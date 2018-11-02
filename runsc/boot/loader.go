// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package boot loads the kernel and runs a container.
package boot

import (
	"fmt"
	mrand "math/rand"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"
	gtime "time"

	specs "github.com/opencontainers/runtime-spec/specs-go"
	"gvisor.googlesource.com/gvisor/pkg/abi/linux"
	"gvisor.googlesource.com/gvisor/pkg/cpuid"
	"gvisor.googlesource.com/gvisor/pkg/log"
	"gvisor.googlesource.com/gvisor/pkg/rand"
	"gvisor.googlesource.com/gvisor/pkg/sentry/arch"
	"gvisor.googlesource.com/gvisor/pkg/sentry/control"
	"gvisor.googlesource.com/gvisor/pkg/sentry/fs/host"
	"gvisor.googlesource.com/gvisor/pkg/sentry/inet"
	"gvisor.googlesource.com/gvisor/pkg/sentry/kernel"
	"gvisor.googlesource.com/gvisor/pkg/sentry/kernel/auth"
	"gvisor.googlesource.com/gvisor/pkg/sentry/loader"
	"gvisor.googlesource.com/gvisor/pkg/sentry/platform"
	"gvisor.googlesource.com/gvisor/pkg/sentry/platform/kvm"
	"gvisor.googlesource.com/gvisor/pkg/sentry/platform/ptrace"
	"gvisor.googlesource.com/gvisor/pkg/sentry/sighandling"
	slinux "gvisor.googlesource.com/gvisor/pkg/sentry/syscalls/linux"
	"gvisor.googlesource.com/gvisor/pkg/sentry/time"
	"gvisor.googlesource.com/gvisor/pkg/sentry/usage"
	"gvisor.googlesource.com/gvisor/pkg/sentry/watchdog"
	"gvisor.googlesource.com/gvisor/pkg/tcpip"
	"gvisor.googlesource.com/gvisor/pkg/tcpip/link/sniffer"
	"gvisor.googlesource.com/gvisor/pkg/tcpip/network/arp"
	"gvisor.googlesource.com/gvisor/pkg/tcpip/network/ipv4"
	"gvisor.googlesource.com/gvisor/pkg/tcpip/network/ipv6"
	"gvisor.googlesource.com/gvisor/pkg/tcpip/stack"
	"gvisor.googlesource.com/gvisor/pkg/tcpip/transport/ping"
	"gvisor.googlesource.com/gvisor/pkg/tcpip/transport/tcp"
	"gvisor.googlesource.com/gvisor/pkg/tcpip/transport/udp"
	"gvisor.googlesource.com/gvisor/runsc/boot/filter"
	"gvisor.googlesource.com/gvisor/runsc/specutils"

	// Include supported socket providers.
	"gvisor.googlesource.com/gvisor/pkg/sentry/socket/epsocket"
	"gvisor.googlesource.com/gvisor/pkg/sentry/socket/hostinet"
	_ "gvisor.googlesource.com/gvisor/pkg/sentry/socket/netlink"
	_ "gvisor.googlesource.com/gvisor/pkg/sentry/socket/netlink/route"
	_ "gvisor.googlesource.com/gvisor/pkg/sentry/socket/unix"
)

// Loader keeps state needed to start the kernel and run the container..
type Loader struct {
	// k is the kernel.
	k *kernel.Kernel

	// ctrl is the control server.
	ctrl *controller

	conf *Config

	// console is set to true if terminal is enabled.
	console bool

	watchdog *watchdog.Watchdog

	// stdioFDs contains stdin, stdout, and stderr.
	stdioFDs []int

	// goferFDs are the FDs that attach the sandbox to the gofers.
	goferFDs []int

	// spec is the base configuration for the root container.
	spec *specs.Spec

	// startSignalForwarding enables forwarding of signals to the sandboxed
	// container. It should be called after the init process is loaded.
	startSignalForwarding func() func()

	// stopSignalForwarding disables forwarding of signals to the sandboxed
	// container. It should be called when a sandbox is destroyed.
	stopSignalForwarding func()

	// restore is set to true if we are restoring a container.
	restore bool

	// rootProcArgs refers to the root sandbox init task.
	rootProcArgs kernel.CreateProcessArgs

	// sandboxID is the ID for the whole sandbox.
	sandboxID string

	// mu guards processes.
	mu sync.Mutex

	// processes maps containers init process and invocation of exec. Root
	// processes are keyed with container ID and pid=0, while exec invocations
	// have the corresponding pid set.
	//
	// processes is guardded by mu.
	processes map[execID]*execProcess
}

// execID uniquely identifies a sentry process.
type execID struct {
	cid string
	pid kernel.ThreadID
}

// execProcess contains the thread group and host TTY of a sentry process.
type execProcess struct {
	tg *kernel.ThreadGroup

	// tty will be nil if the process is not attached to a terminal.
	tty *host.TTYFileOperations
}

func init() {
	// Initialize the random number generator.
	mrand.Seed(gtime.Now().UnixNano())

	// Register the global syscall table.
	kernel.RegisterSyscallTable(slinux.AMD64)
}

// Args are the arguments for New().
type Args struct {
	// Id is the sandbox ID.
	ID string
	// Spec is the sandbox specification.
	Spec *specs.Spec
	// Conf is the system configuration.
	Conf *Config
	// ControllerFD is the FD to the URPC controller.
	ControllerFD int
	// DeviceFD is an optional argument that is passed to the platform.
	DeviceFD int
	// GoferFDs is an array of FDs used to connect with the Gofer.
	GoferFDs []int
	// StdioFDs is the stdio for the application.
	StdioFDs []int
	// Console is set to true if using TTY.
	Console bool
	// NumCPU is the number of CPUs to create inside the sandbox.
	NumCPU int
	// TotalMem is the initial amount of total memory to report back to the
	// container.
	TotalMem uint64
	// UserLogFD is the file descriptor to write user logs to.
	UserLogFD int
}

// New initializes a new kernel loader configured by spec.
// New also handles setting up a kernel for restoring a container.
func New(args Args) (*Loader, error) {
	// We initialize the rand package now to make sure /dev/urandom is pre-opened
	// on kernels that do not support getrandom(2).
	if err := rand.Init(); err != nil {
		return nil, fmt.Errorf("error setting up rand: %v", err)
	}

	if err := usage.Init(); err != nil {
		return nil, fmt.Errorf("error setting up memory usage: %v", err)
	}

	// Create kernel and platform.
	p, err := createPlatform(args.Conf, args.DeviceFD)
	if err != nil {
		return nil, fmt.Errorf("error creating platform: %v", err)
	}
	k := &kernel.Kernel{
		Platform: p,
	}

	// Create VDSO.
	//
	// Pass k as the platform since it is savable, unlike the actual platform.
	vdso, err := loader.PrepareVDSO(k)
	if err != nil {
		return nil, fmt.Errorf("error creating vdso: %v", err)
	}

	// Create timekeeper.
	tk, err := kernel.NewTimekeeper(k, vdso.ParamPage.FileRange())
	if err != nil {
		return nil, fmt.Errorf("error creating timekeeper: %v", err)
	}
	tk.SetClocks(time.NewCalibratedClocks())

	if err := enableStrace(args.Conf); err != nil {
		return nil, fmt.Errorf("failed to enable strace: %v", err)
	}

	// Create an empty network stack because the network namespace may be empty at
	// this point. Netns is configured before Run() is called. Netstack is
	// configured using a control uRPC message. Host network is configured inside
	// Run().
	networkStack, err := newEmptyNetworkStack(args.Conf, k)
	if err != nil {
		return nil, fmt.Errorf("failed to create network: %v", err)
	}

	// Create capabilities.
	caps, err := specutils.Capabilities(args.Spec.Process.Capabilities)
	if err != nil {
		return nil, fmt.Errorf("error creating capabilities: %v", err)
	}

	// Convert the spec's additional GIDs to KGIDs.
	extraKGIDs := make([]auth.KGID, 0, len(args.Spec.Process.User.AdditionalGids))
	for _, GID := range args.Spec.Process.User.AdditionalGids {
		extraKGIDs = append(extraKGIDs, auth.KGID(GID))
	}

	// Create credentials.
	creds := auth.NewUserCredentials(
		auth.KUID(args.Spec.Process.User.UID),
		auth.KGID(args.Spec.Process.User.GID),
		extraKGIDs,
		caps,
		auth.NewRootUserNamespace())

	if args.NumCPU == 0 {
		args.NumCPU = runtime.NumCPU()
	}
	log.Infof("CPUs: %d", args.NumCPU)

	if args.TotalMem > 0 {
		// Adjust the total memory returned by the Sentry so that applications that
		// use /proc/meminfo can make allocations based on this limit.
		usage.MinimumTotalMemoryBytes = args.TotalMem
		log.Infof("Setting total memory to %.2f GB", float64(args.TotalMem)/(2^30))
	}

	// Initiate the Kernel object, which is required by the Context passed
	// to createVFS in order to mount (among other things) procfs.
	if err = k.Init(kernel.InitKernelArgs{
		FeatureSet:                  cpuid.HostFeatureSet(),
		Timekeeper:                  tk,
		RootUserNamespace:           creds.UserNamespace,
		NetworkStack:                networkStack,
		ApplicationCores:            uint(args.NumCPU),
		Vdso:                        vdso,
		RootUTSNamespace:            kernel.NewUTSNamespace(args.Spec.Hostname, "", creds.UserNamespace),
		RootIPCNamespace:            kernel.NewIPCNamespace(creds.UserNamespace),
		RootAbstractSocketNamespace: kernel.NewAbstractSocketNamespace(),
	}); err != nil {
		return nil, fmt.Errorf("error initializing kernel: %v", err)
	}

	// Turn on packet logging if enabled.
	if args.Conf.LogPackets {
		log.Infof("Packet logging enabled")
		atomic.StoreUint32(&sniffer.LogPackets, 1)
	} else {
		log.Infof("Packet logging disabled")
		atomic.StoreUint32(&sniffer.LogPackets, 0)
	}

	// Create a watchdog.
	watchdog := watchdog.New(k, watchdog.DefaultTimeout, args.Conf.WatchdogAction)

	// Create the control server using the provided FD.
	//
	// This must be done *after* we have initialized the kernel since the
	// controller is used to configure the kernel's network stack.
	//
	// This should also be *before* we create the process, since a
	// misconfigured process will cause an error, and we want the control
	// server up before that so that we don't time out trying to connect to
	// it.
	ctrl, err := newController(args.ControllerFD, k, watchdog)
	if err != nil {
		return nil, fmt.Errorf("error creating control server: %v", err)
	}

	procArgs, err := newProcess(args.ID, args.Spec, creds, k)
	if err != nil {
		return nil, fmt.Errorf("failed to create init process for root container: %v", err)
	}

	if err := initCompatLogs(args.UserLogFD); err != nil {
		return nil, fmt.Errorf("init compat logs: %v", err)
	}

	l := &Loader{
		k:            k,
		ctrl:         ctrl,
		conf:         args.Conf,
		console:      args.Console,
		watchdog:     watchdog,
		spec:         args.Spec,
		goferFDs:     args.GoferFDs,
		stdioFDs:     args.StdioFDs,
		rootProcArgs: procArgs,
		sandboxID:    args.ID,
		processes:    make(map[execID]*execProcess),
	}

	// We don't care about child signals; some platforms can generate a
	// tremendous number of useless ones (I'm looking at you, ptrace).
	if err := sighandling.IgnoreChildStop(); err != nil {
		return nil, fmt.Errorf("failed to ignore child stop signals: %v", err)
	}

	// Handle signals by forwarding them to the root container process
	// (except for panic signal, which should cause a panic).
	l.startSignalForwarding = sighandling.PrepareHandler(func(sig linux.Signal) {
		// Panic signal should cause a panic.
		if args.Conf.PanicSignal != -1 && sig == linux.Signal(args.Conf.PanicSignal) {
			panic("Signal-induced panic")
		}

		// Otherwise forward to root container.
		deliveryMode := DeliverToProcess
		if args.Console {
			// Since we are running with a console, we should
			// forward the signal to the foreground process group
			// so that job control signals like ^C can be handled
			// properly.
			deliveryMode = DeliverToForegroundProcessGroup
		}
		if err := l.signal(args.ID, 0, int32(sig), deliveryMode); err != nil {
			log.Warningf("error sending signal %v to container %q: %v", sig, args.ID, err)
		}
	})

	ctrl.manager.l = l
	return l, nil
}

// newProcess creates a process that can be run with kernel.CreateProcess.
func newProcess(id string, spec *specs.Spec, creds *auth.Credentials, k *kernel.Kernel) (kernel.CreateProcessArgs, error) {
	// Create initial limits.
	ls, err := createLimitSet(spec)
	if err != nil {
		return kernel.CreateProcessArgs{}, fmt.Errorf("error creating limits: %v", err)
	}

	// Create the process arguments.
	procArgs := kernel.CreateProcessArgs{
		Argv:                    spec.Process.Args,
		Envv:                    spec.Process.Env,
		WorkingDirectory:        spec.Process.Cwd, // Defaults to '/' if empty.
		Credentials:             creds,
		Umask:                   0022,
		Limits:                  ls,
		MaxSymlinkTraversals:    linux.MaxSymlinkTraversals,
		UTSNamespace:            k.RootUTSNamespace(),
		IPCNamespace:            k.RootIPCNamespace(),
		AbstractSocketNamespace: k.RootAbstractSocketNamespace(),
		ContainerID:             id,
	}
	return procArgs, nil
}

// Destroy cleans up all resources used by the loader.
//
// Note that this will block until all open control server connections have
// been closed. For that reason, this should NOT be called in a defer, because
// a panic in a control server rpc would then hang forever.
func (l *Loader) Destroy() {
	if l.ctrl != nil {
		l.ctrl.srv.Stop()
	}
	if l.stopSignalForwarding != nil {
		l.stopSignalForwarding()
	}
	l.watchdog.Stop()
}

func createPlatform(conf *Config, deviceFD int) (platform.Platform, error) {
	switch conf.Platform {
	case PlatformPtrace:
		log.Infof("Platform: ptrace")
		return ptrace.New()
	case PlatformKVM:
		log.Infof("Platform: kvm")
		if deviceFD < 0 {
			return nil, fmt.Errorf("kvm device FD must be provided")
		}
		return kvm.New(os.NewFile(uintptr(deviceFD), "kvm device"))
	default:
		return nil, fmt.Errorf("invalid platform %v", conf.Platform)
	}
}

// Run runs the root container..
func (l *Loader) Run() error {
	err := l.run()
	l.ctrl.manager.startResultChan <- err
	if err != nil {
		// Give the controller some time to send the error to the
		// runtime. If we return too quickly here the process will exit
		// and the control connection will be closed before the error
		// is returned.
		gtime.Sleep(2 * gtime.Second)
		return err
	}
	return nil
}

func (l *Loader) run() error {
	if l.conf.Network == NetworkHost {
		// Delay host network configuration to this point because network namespace
		// is configured after the loader is created and before Run() is called.
		log.Debugf("Configuring host network")
		stack := l.k.NetworkStack().(*hostinet.Stack)
		if err := stack.Configure(); err != nil {
			return err
		}
	}

	// Finally done with all configuration. Setup filters before user code
	// is loaded.
	if l.conf.DisableSeccomp {
		filter.Report("syscall filter is DISABLED. Running in less secure mode.")
	} else {
		opts := filter.Options{
			Platform:     l.k.Platform,
			HostNetwork:  l.conf.Network == NetworkHost,
			ControllerFD: l.ctrl.srv.FD(),
		}
		if err := filter.Install(opts); err != nil {
			return fmt.Errorf("Failed to install seccomp filters: %v", err)
		}
	}

	// If we are restoring, we do not want to create a process.
	// l.restore is set by the container manager when a restore call is made.
	if !l.restore {
		if err := setupContainerFS(
			&l.rootProcArgs,
			l.spec,
			l.conf,
			l.stdioFDs,
			l.goferFDs,
			l.console,
			l.rootProcArgs.Credentials,
			l.rootProcArgs.Limits,
			l.k,
			"" /* CID, which isn't needed for the root container */); err != nil {
			return err
		}

		rootCtx := l.rootProcArgs.NewContext(l.k)
		rootMns := l.k.RootMountNamespace()
		if err := setExecutablePath(rootCtx, rootMns, &l.rootProcArgs); err != nil {
			return fmt.Errorf("error setting executable path for %+v: %v", l.rootProcArgs, err)
		}

		// Create the root container init task.
		_, _, err := l.k.CreateProcess(l.rootProcArgs)
		if err != nil {
			return fmt.Errorf("failed to create init process: %v", err)
		}

		// CreateProcess takes a reference on FDMap if successful.
		l.rootProcArgs.FDMap.DecRef()
	}

	eid := execID{cid: l.sandboxID}
	ep := execProcess{tg: l.k.GlobalInit()}
	if l.console {
		ttyFile := l.rootProcArgs.FDMap.GetFile(0)
		defer ttyFile.DecRef()
		ep.tty = ttyFile.FileOperations.(*host.TTYFileOperations)
	}
	l.mu.Lock()
	l.processes[eid] = &ep
	l.mu.Unlock()

	// Start signal forwarding only after an init process is created.
	l.stopSignalForwarding = l.startSignalForwarding()

	log.Infof("Process should have started...")
	l.watchdog.Start()
	return l.k.Start()
}

// startContainer starts a child container. It returns the thread group ID of
// the newly created process.
func (l *Loader) startContainer(k *kernel.Kernel, spec *specs.Spec, conf *Config, cid string, files []*os.File) error {
	// Create capabilities.
	caps, err := specutils.Capabilities(spec.Process.Capabilities)
	if err != nil {
		return fmt.Errorf("error creating capabilities: %v", err)
	}

	// Convert the spec's additional GIDs to KGIDs.
	extraKGIDs := make([]auth.KGID, 0, len(spec.Process.User.AdditionalGids))
	for _, GID := range spec.Process.User.AdditionalGids {
		extraKGIDs = append(extraKGIDs, auth.KGID(GID))
	}

	// Create credentials. We reuse the root user namespace because the
	// sentry currently supports only 1 mount namespace, which is tied to a
	// single user namespace. Thus we must run in the same user namespace
	// to access mounts.
	// TODO: Create a new mount namespace for the container.
	creds := auth.NewUserCredentials(
		auth.KUID(spec.Process.User.UID),
		auth.KGID(spec.Process.User.GID),
		extraKGIDs,
		caps,
		l.k.RootUserNamespace())

	procArgs, err := newProcess(cid, spec, creds, l.k)
	if err != nil {
		return fmt.Errorf("failed to create new process: %v", err)
	}

	// Can't take ownership away from os.File. dup them to get a new FDs.
	var ioFDs []int
	for _, f := range files {
		fd, err := syscall.Dup(int(f.Fd()))
		if err != nil {
			return fmt.Errorf("failed to dup file: %v", err)
		}
		f.Close()
		ioFDs = append(ioFDs, fd)
	}

	stdioFDs := ioFDs[:3]
	goferFDs := ioFDs[3:]
	if err := setupContainerFS(
		&procArgs,
		spec,
		conf,
		stdioFDs,
		goferFDs,
		false,
		creds,
		procArgs.Limits,
		k,
		cid); err != nil {
		return fmt.Errorf("failed to create new process: %v", err)
	}

	// setFileSystemForProcess dup'd stdioFDs, so we can close them.
	for i, fd := range stdioFDs {
		if err := syscall.Close(fd); err != nil {
			return fmt.Errorf("failed to close stdioFD #%d: %v", i, fd)
		}
	}

	ctx := procArgs.NewContext(l.k)
	mns := k.RootMountNamespace()
	if err := setExecutablePath(ctx, mns, &procArgs); err != nil {
		return fmt.Errorf("error setting executable path for %+v: %v", procArgs, err)
	}

	tg, _, err := l.k.CreateProcess(procArgs)
	if err != nil {
		return fmt.Errorf("failed to create process in sentry: %v", err)
	}

	// CreateProcess takes a reference on FDMap if successful.
	procArgs.FDMap.DecRef()

	l.mu.Lock()
	defer l.mu.Unlock()
	eid := execID{cid: cid}
	l.processes[eid] = &execProcess{tg: tg}

	return nil
}

// destroyContainer stops a container if it is still running and cleans up its
// filesystem.
func (l *Loader) destroyContainer(cid string) error {
	// First kill and wait for all processes in the container.
	if err := l.signal(cid, 0, int32(linux.SIGKILL), DeliverToAllProcesses); err != nil {
		return fmt.Errorf("failed to SIGKILL all container processes: %v", err)
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	// Remove all container thread groups from the map.
	for key := range l.processes {
		if key.cid == cid {
			delete(l.processes, key)
		}
	}

	ctx := l.rootProcArgs.NewContext(l.k)
	if err := destroyContainerFS(ctx, cid, l.k); err != nil {
		return fmt.Errorf("failed to destroy filesystem for container %q: %v", cid, err)
	}

	// We made it!
	log.Debugf("Container destroyed %q", cid)
	return nil
}

func (l *Loader) executeAsync(args *control.ExecArgs) (kernel.ThreadID, error) {
	// Get the container Root Dirent from the Task, since we must run this
	// process with the same Root.
	l.mu.Lock()
	rootKey := execID{cid: args.ContainerID}
	ep, ok := l.processes[rootKey]
	l.mu.Unlock()
	if !ok {
		return 0, fmt.Errorf("no such container: %q", args.ContainerID)
	}
	ep.tg.Leader().WithMuLocked(func(t *kernel.Task) {
		args.Root = t.FSContext().RootDirectory()
	})
	if args.Root != nil {
		defer args.Root.DecRef()
	}

	// Start the process.
	proc := control.Proc{Kernel: l.k}
	tg, tgid, ttyFile, err := control.ExecAsync(&proc, args)
	if err != nil {
		return 0, err
	}

	// Insert the process into processes so that we can wait on it
	// later.
	l.mu.Lock()
	defer l.mu.Unlock()
	eid := execID{cid: args.ContainerID, pid: tgid}
	l.processes[eid] = &execProcess{
		tg:  tg,
		tty: ttyFile,
	}
	log.Debugf("updated processes: %v", l.processes)

	return tgid, nil
}

// waitContainer waits for the init process of a container to exit.
func (l *Loader) waitContainer(cid string, waitStatus *uint32) error {
	// Don't defer unlock, as doing so would make it impossible for
	// multiple clients to wait on the same container.
	l.mu.Lock()
	eid := execID{cid: cid}
	ep, ok := l.processes[eid]
	l.mu.Unlock()
	if !ok {
		return fmt.Errorf("can't find process for container %q in %v", cid, l.processes)
	}

	// If the thread either has already exited or exits during waiting,
	// consider the container exited.
	ws := l.wait(ep.tg)
	*waitStatus = ws
	return nil
}

func (l *Loader) waitPID(tgid kernel.ThreadID, cid string, clearStatus bool, waitStatus *uint32) error {
	// If the process was started via runsc exec, it will have an
	// entry in l.processes.
	l.mu.Lock()
	eid := execID{cid: cid, pid: tgid}
	ep, ok := l.processes[eid]
	l.mu.Unlock()
	if ok {
		ws := l.wait(ep.tg)
		*waitStatus = ws
		if clearStatus {
			// Remove tg from the cache.
			l.mu.Lock()
			delete(l.processes, eid)
			log.Debugf("updated processes (removal): %v", l.processes)
			l.mu.Unlock()
		}
		return nil
	}

	// This process wasn't created by runsc exec or start, so just find it
	// by PID and hope it hasn't exited yet.
	tg := l.k.TaskSet().Root.ThreadGroupWithID(kernel.ThreadID(tgid))
	if tg == nil {
		return fmt.Errorf("no thread group with ID %d", tgid)
	}
	ws := l.wait(tg)
	*waitStatus = ws
	return nil
}

// wait waits for the process with TGID 'tgid' in a container's PID namespace
// to exit.
func (l *Loader) wait(tg *kernel.ThreadGroup) uint32 {
	tg.WaitExited()
	return tg.ExitStatus().Status()
}

// WaitForStartSignal waits for a start signal from the control server.
func (l *Loader) WaitForStartSignal() {
	<-l.ctrl.manager.startChan
}

// NotifyLoaderCreated sends a signal to the container manager that this
// loader has been created.
func (l *Loader) NotifyLoaderCreated() {
	l.ctrl.manager.loaderCreatedChan <- struct{}{}
}

// WaitExit waits for the root container to exit, and returns its exit status.
func (l *Loader) WaitExit() kernel.ExitStatus {
	// Wait for container.
	l.k.WaitExited()

	return l.k.GlobalInit().ExitStatus()
}

func newEmptyNetworkStack(conf *Config, clock tcpip.Clock) (inet.Stack, error) {
	switch conf.Network {
	case NetworkHost:
		return hostinet.NewStack(), nil

	case NetworkNone, NetworkSandbox:
		// NetworkNone sets up loopback using netstack.
		netProtos := []string{ipv4.ProtocolName, ipv6.ProtocolName, arp.ProtocolName}
		protoNames := []string{tcp.ProtocolName, udp.ProtocolName, ping.ProtocolName4}
		s := epsocket.Stack{stack.New(netProtos, protoNames, stack.Options{
			Clock: clock,
			Stats: epsocket.Metrics,
		})}
		if err := s.Stack.SetTransportProtocolOption(tcp.ProtocolNumber, tcp.SACKEnabled(true)); err != nil {
			return nil, fmt.Errorf("failed to enable SACK: %v", err)
		}
		return &s, nil

	default:
		panic(fmt.Sprintf("invalid network configuration: %v", conf.Network))
	}
}

// signal sends a signal to one or more processes in a container. If PID is 0,
// then the container init process is used. Depending on the SignalDeliveryMode
// option, the signal may be sent directly to the indicated process, to all
// processes in the container, or to the foreground process group.
func (l *Loader) signal(cid string, pid, signo int32, mode SignalDeliveryMode) error {
	if pid < 0 {
		return fmt.Errorf("failed to signal container %q PID %d: PID must be positive", cid, pid)
	}

	eid := execID{
		cid: cid,
		pid: kernel.ThreadID(pid),
	}
	l.mu.Lock()
	ep, ok := l.processes[eid]
	l.mu.Unlock()

	switch mode {
	case DeliverToProcess:
		if ok {
			// Send signal directly to the identified process.
			return ep.tg.SendSignal(&arch.SignalInfo{Signo: signo})
		}

		// The caller may be signaling a process not started directly via exec.
		// In this case, find the process in the container's PID namespace and
		// signal it.
		ep, ok := l.processes[execID{cid: cid}]
		if !ok {
			return fmt.Errorf("no container with ID: %q", cid)
		}
		tg := ep.tg.PIDNamespace().ThreadGroupWithID(kernel.ThreadID(pid))
		if tg == nil {
			return fmt.Errorf("failed to signal container %q PID %d: no such process", cid, pid)
		}
		if tg.Leader().ContainerID() != cid {
			return fmt.Errorf("process %d is part of a different container: %q", pid, tg.Leader().ContainerID())
		}
		return tg.SendSignal(&arch.SignalInfo{Signo: signo})

	case DeliverToForegroundProcessGroup:
		if !ok {
			return fmt.Errorf("failed to signal foreground process group for container %q PID %d: no such PID", cid, pid)
		}

		// Lookup foreground process group from the TTY for the given process,
		// and send the signal to it.
		if ep.tty == nil {
			return fmt.Errorf("failed to signal foreground process group in container %q PID %d: no TTY attached", cid, pid)
		}
		pg := ep.tty.ForegroundProcessGroup()
		if pg == nil {
			// No foreground process group has been set. Signal the
			// original thread group.
			log.Warningf("No foreground process group for container %q and PID %d. Sending signal directly to PID %d.", cid, pid, pid)
			return ep.tg.SendSignal(&arch.SignalInfo{Signo: signo})
		}
		// Send the signal to all processes in the process group.
		var lastErr error
		for _, tg := range l.k.TaskSet().Root.ThreadGroups() {
			if tg.ProcessGroup() != pg {
				continue
			}
			if err := tg.SendSignal(&arch.SignalInfo{Signo: signo}); err != nil {
				lastErr = err
			}
		}
		return lastErr
	case DeliverToAllProcesses:
		if !ok {
			return fmt.Errorf("failed to signal all processes in container %q PID %d: no such PID", cid, pid)
		}

		// Pause the kernel to prevent new processes from being created while
		// the signal is delivered. This prevents process leaks when SIGKILL is
		// sent to the entire container.
		l.k.Pause()
		if err := l.k.SendContainerSignal(cid, &arch.SignalInfo{Signo: signo}); err != nil {
			l.k.Unpause()
			return err
		}
		l.k.Unpause()

		// If SIGKILLing all processes, wait for them to exit.
		if linux.Signal(signo) == linux.SIGKILL {
			for _, t := range l.k.TaskSet().Root.Tasks() {
				if t.ContainerID() == cid {
					t.ThreadGroup().WaitExited()
				}
			}
		}
		return nil
	default:
		panic(fmt.Sprintf("unknown signal signal delivery mode %v", mode))
	}
}
