// Copyright 2018 Google Inc.
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

// Package sandbox creates and manipulates sandboxes.
package sandbox

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"time"

	specs "github.com/opencontainers/runtime-spec/specs-go"
	"golang.org/x/sys/unix"
	"gvisor.googlesource.com/gvisor/pkg/control/client"
	"gvisor.googlesource.com/gvisor/pkg/control/server"
	"gvisor.googlesource.com/gvisor/pkg/log"
	"gvisor.googlesource.com/gvisor/pkg/sentry/control"
	"gvisor.googlesource.com/gvisor/pkg/urpc"
	"gvisor.googlesource.com/gvisor/runsc/boot"
	"gvisor.googlesource.com/gvisor/runsc/specutils"
)

// Sandbox wraps a sandbox process.
//
// It is used to start/stop sandbox process (and associated processes like
// gofers), as well as for running and manipulating containers inside a running
// sandbox.
type Sandbox struct {
	// ID is the id of the sandbox. By convention, this is the same ID as
	// the first container run in the sandbox.
	ID string `json:"id"`

	// Pid is the pid of the running sandbox. May be 0 is the sandbox is
	// not running.
	Pid int `json:"pid"`

	// GoferPid is the pid of the gofer running along side the sandbox. May
	// be 0 if the gofer has been killed or it's not being used.
	GoferPid int `json:"goferPid"`
}

// Create creates the sandbox process.
//
// If restoreFile is not empty, the sandbox will be restored from file.
func Create(id string, spec *specs.Spec, conf *boot.Config, bundleDir, consoleSocket string, restoreFile string) (*Sandbox, error) {
	s := &Sandbox{ID: id}

	binPath, err := specutils.BinPath()
	if err != nil {
		return nil, err
	}

	// Create the gofer process.
	ioFiles, err := s.createGoferProcess(spec, conf, bundleDir, binPath)
	if err != nil {
		return nil, err
	}

	// Create the sandbox process.
	if err := s.createSandboxProcess(spec, conf, bundleDir, consoleSocket, binPath, ioFiles, restoreFile); err != nil {
		return nil, err
	}

	// Wait for the control server to come up (or timeout).
	if err := s.waitForCreated(10 * time.Second); err != nil {
		return nil, err
	}

	return s, nil
}

// StartRoot starts running the root container process inside the sandbox.
func (s *Sandbox) StartRoot(spec *specs.Spec, conf *boot.Config) error {
	log.Debugf("Start root sandbox %q, pid: %d", s.ID, s.Pid)
	conn, err := s.connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Configure the network.
	if err := setupNetwork(conn, s.Pid, spec, conf); err != nil {
		return fmt.Errorf("error setting up network: %v", err)
	}

	// Send a message to the sandbox control server to start the root
	// container.
	if err := conn.Call(boot.RootContainerStart, &s.ID, nil); err != nil {
		return fmt.Errorf("error starting root container %v: %v", spec.Process.Args, err)
	}

	return nil
}

// Start starts running a non-root container inside the sandbox.
func (s *Sandbox) Start(spec *specs.Spec, conf *boot.Config, cid string) error {
	log.Debugf("Start non-root container sandbox %q, pid: %d", s.ID, s.Pid)
	conn, err := s.connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	args := boot.StartArgs{
		Spec: spec,
		Conf: conf,
		CID:  cid,
	}
	if err := conn.Call(boot.ContainerStart, args, nil); err != nil {
		return fmt.Errorf("error starting non-root container %v: %v", spec.Process.Args, err)
	}

	return nil
}

// Processes retrieves the list of processes and associated metadata for a
// given container in this sandbox.
func (s *Sandbox) Processes(cid string) ([]*control.Process, error) {
	log.Debugf("Getting processes for container %q in sandbox %q", cid, s.ID)
	conn, err := s.connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var pl []*control.Process
	// TODO: Pass in the container id (cid) here. The sandbox
	// should return process info for only that container.
	if err := conn.Call(boot.ContainerProcesses, nil, &pl); err != nil {
		return nil, fmt.Errorf("error retrieving process data from sandbox: %v", err)
	}
	return pl, nil
}

// Execute runs the specified command in the container.
func (s *Sandbox) Execute(cid string, e *control.ExecArgs) (syscall.WaitStatus, error) {
	log.Debugf("Executing new process in container %q in sandbox %q", cid, s.ID)
	conn, err := s.connect()
	if err != nil {
		return 0, s.connError(err)
	}
	defer conn.Close()

	// Send a message to the sandbox control server to start the container.
	var waitStatus uint32
	// TODO: Pass in the container id (cid) here. The sandbox
	// should execute in the context of that container.
	if err := conn.Call(boot.ContainerExecute, e, &waitStatus); err != nil {
		return 0, fmt.Errorf("error executing in sandbox: %v", err)
	}

	return syscall.WaitStatus(waitStatus), nil
}

// Event retrieves stats about the sandbox such as memory and CPU utilization.
func (s *Sandbox) Event(cid string) (*boot.Event, error) {
	log.Debugf("Getting events for container %q in sandbox %q", cid, s.ID)
	conn, err := s.connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var e boot.Event
	// TODO: Pass in the container id (cid) here. The sandbox
	// should return events only for that container.
	if err := conn.Call(boot.ContainerEvent, nil, &e); err != nil {
		return nil, fmt.Errorf("error retrieving event data from sandbox: %v", err)
	}
	e.ID = cid
	return &e, nil
}

func (s *Sandbox) connect() (*urpc.Client, error) {
	log.Debugf("Connecting to sandbox %q", s.ID)
	conn, err := client.ConnectTo(boot.ControlSocketAddr(s.ID))
	if err != nil {
		return nil, s.connError(err)
	}
	return conn, nil
}

func (s *Sandbox) connError(err error) error {
	return fmt.Errorf("error connecting to control server at pid %d: %v", s.Pid, err)
}

func (s *Sandbox) createGoferProcess(spec *specs.Spec, conf *boot.Config, bundleDir, binPath string) ([]*os.File, error) {
	if conf.FileAccess != boot.FileAccessProxy {
		// Don't start a gofer. The sandbox will access host FS directly.
		return nil, nil
	}

	// Start with the general config flags.
	args := conf.ToFlags()
	args = append(args, "gofer", "--bundle", bundleDir)

	// Add root mount and then add any other additional mounts.
	mountCount := 1
	for _, m := range spec.Mounts {
		if specutils.Is9PMount(m) {
			mountCount++
		}
	}

	sandEnds := make([]*os.File, 0, mountCount)
	goferEnds := make([]*os.File, 0, mountCount)
	for i := 0; i < mountCount; i++ {
		// Create socket that connects the sandbox and gofer.
		fds, err := syscall.Socketpair(syscall.AF_UNIX, syscall.SOCK_STREAM|syscall.SOCK_CLOEXEC, 0)
		if err != nil {
			return nil, err
		}
		sandEnds = append(sandEnds, os.NewFile(uintptr(fds[0]), "sandbox io fd"))

		goferEnd := os.NewFile(uintptr(fds[1]), "gofer io fd")
		defer goferEnd.Close()
		goferEnds = append(goferEnds, goferEnd)

		args = append(args, fmt.Sprintf("--io-fds=%d", 3+i))
	}

	cmd := exec.Command(binPath, args...)
	cmd.ExtraFiles = goferEnds

	// Setup any uid/gid mappings, and create or join the configured user
	// namespace so the gofer's view of the filesystem aligns with the
	// users in the sandbox.
	setUIDGIDMappings(cmd, spec)
	nss := filterNS([]specs.LinuxNamespaceType{specs.UserNamespace}, spec)

	// Start the gofer in the given namespace.
	log.Debugf("Starting gofer: %s %v", binPath, args)
	if err := startInNS(cmd, nss); err != nil {
		return nil, err
	}
	s.GoferPid = cmd.Process.Pid
	log.Infof("Gofer started, pid: %d", cmd.Process.Pid)
	return sandEnds, nil
}

// createSandboxProcess starts the sandbox as a subprocess by running the "boot"
// command, passing in the bundle dir.
func (s *Sandbox) createSandboxProcess(spec *specs.Spec, conf *boot.Config, bundleDir, consoleSocket, binPath string, ioFiles []*os.File, restoreFile string) error {
	// nextFD is used to get unused FDs that we can pass to the sandbox.  It
	// starts at 3 because 0, 1, and 2 are taken by stdin/out/err.
	nextFD := 3

	// Create control server socket here and donate FD to child process because
	// it may be in a different network namespace and won't be reachable from
	// outside.
	fd, err := server.CreateSocket(boot.ControlSocketAddr(s.ID))
	if err != nil {
		return fmt.Errorf("error creating control server socket for sandbox %q: %v", s.ID, err)
	}

	consoleEnabled := consoleSocket != ""

	cmd := exec.Command(binPath, conf.ToFlags()...)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.Args = append(cmd.Args,
		"boot",
		"--bundle", bundleDir,
		"--controller-fd="+strconv.Itoa(nextFD),
		fmt.Sprintf("--console=%t", consoleEnabled))

	controllerFile := os.NewFile(uintptr(fd), "control_server_socket")
	defer controllerFile.Close()
	cmd.ExtraFiles = append(cmd.ExtraFiles, controllerFile)

	// If a restore filename was given, open the file and append its FD to Args
	// and the file to ExtraFiles.
	if restoreFile != "" {
		// Create the image file and open for reading.
		rF, err := os.Open(restoreFile)
		if err != nil {
			return fmt.Errorf("os.Open(%q) failed: %v", restoreFile, err)
		}
		defer rF.Close()

		nextFD++
		cmd.Args = append(cmd.Args, "--restore-fd="+strconv.Itoa(nextFD))
		cmd.ExtraFiles = append(cmd.ExtraFiles, rF)
	}
	nextFD++

	// If there is a gofer, sends all socket ends to the sandbox.
	for _, f := range ioFiles {
		defer f.Close()
		cmd.ExtraFiles = append(cmd.ExtraFiles, f)
		cmd.Args = append(cmd.Args, "--io-fds="+strconv.Itoa(nextFD))
		nextFD++
	}

	// If the console control socket file is provided, then create a new
	// pty master/slave pair and set the tty on the sandbox process.
	if consoleEnabled {
		// setupConsole will send the master on the socket, and return
		// the slave.
		tty, err := setupConsole(consoleSocket)
		if err != nil {
			return fmt.Errorf("error setting up control socket %q: %v", consoleSocket, err)
		}
		defer tty.Close()

		cmd.Stdin = tty
		cmd.Stdout = tty
		cmd.Stderr = tty
		cmd.SysProcAttr.Setctty = true
		cmd.SysProcAttr.Ctty = int(tty.Fd())
	} else {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	// Detach from this session, otherwise cmd will get SIGHUP and SIGCONT
	// when re-parented.
	cmd.SysProcAttr.Setsid = true

	// nss is the set of namespaces to join or create before starting the sandbox
	// process. IPC and UTS namespaces from the host are not used as they
	// are virtualized inside the sandbox. Be paranoid and run inside an empty
	// namespace for these.
	log.Infof("Sandbox will be started in new IPC and UTS namespaces")
	nss := []specs.LinuxNamespace{
		{Type: specs.IPCNamespace},
		{Type: specs.UTSNamespace},
	}

	if conf.Platform == boot.PlatformPtrace {
		// TODO: Also set a new PID namespace so that we limit
		// access to other host processes.
		log.Infof("Sandbox will be started in the current PID namespace")
	} else {
		log.Infof("Sandbox will be started in a new PID namespace")
		nss = append(nss, specs.LinuxNamespace{Type: specs.PIDNamespace})
	}

	if conf.FileAccess == boot.FileAccessProxy {
		log.Infof("Sandbox will be started in new mount namespace")
		nss = append(nss, specs.LinuxNamespace{Type: specs.MountNamespace})
	} else {
		log.Infof("Sandbox will be started in the current mount namespace")
	}

	// Joins the network namespace if network is enabled. the sandbox talks
	// directly to the host network, which may have been configured in the
	// namespace.
	if ns, ok := getNS(specs.NetworkNamespace, spec); ok && conf.Network != boot.NetworkNone {
		log.Infof("Sandbox will be started in the container's network namespace: %+v", ns)
		nss = append(nss, ns)
	} else {
		log.Infof("Sandbox will be started in new network namespace")
		nss = append(nss, specs.LinuxNamespace{Type: specs.NetworkNamespace})
	}

	// User namespace depends on the following options:
	//   - Host network/filesystem: requires to run inside the user namespace
	//       specified in the spec or the current namespace if none is configured.
	//   - Gofer: when using a Gofer, the sandbox process can run isolated in an
	//       empty namespace.
	if conf.Network == boot.NetworkHost || conf.FileAccess == boot.FileAccessDirect {
		if userns, ok := getNS(specs.UserNamespace, spec); ok {
			log.Infof("Sandbox will be started in container's user namespace: %+v", userns)
			nss = append(nss, userns)
			setUIDGIDMappings(cmd, spec)
		} else {
			log.Infof("Sandbox will be started in the current user namespace")
		}
		// When running in the caller's defined user namespace, apply the same
		// capabilities to the sandbox process to ensure it abides to the same
		// rules.
		cmd.Args = append(cmd.Args, "--apply-caps=true")

	} else {
		log.Infof("Sandbox will be started in new user namespace")
		nss = append(nss, specs.LinuxNamespace{Type: specs.UserNamespace})
	}

	log.Debugf("Starting sandbox: %s %v", binPath, cmd.Args)
	if err := startInNS(cmd, nss); err != nil {
		return err
	}
	s.Pid = cmd.Process.Pid
	log.Infof("Sandbox started, pid: %d", s.Pid)

	return nil
}

// waitForCreated waits for the sandbox subprocess control server to be
// running, at which point the sandbox is in Created state.
func (s *Sandbox) waitForCreated(timeout time.Duration) error {
	log.Debugf("Waiting for sandbox %q creation", s.ID)

	ready := func() (bool, error) {
		c, err := client.ConnectTo(boot.ControlSocketAddr(s.ID))
		if err != nil {
			return false, nil
		}
		// It's alive!
		c.Close()
		return true, nil
	}
	if err := specutils.WaitForReady(s.Pid, timeout, ready); err != nil {
		return fmt.Errorf("unexpected error waiting for sandbox %q, err: %v", s.ID, err)
	}
	return nil
}

// Wait waits for the containerized process to exit, and returns its WaitStatus.
func (s *Sandbox) Wait(cid string) (syscall.WaitStatus, error) {
	log.Debugf("Waiting for container %q in sandbox %q", cid, s.ID)
	var ws syscall.WaitStatus
	conn, err := s.connect()
	if err != nil {
		return ws, err
	}
	defer conn.Close()

	if err := conn.Call(boot.ContainerWait, &cid, &ws); err != nil {
		return ws, fmt.Errorf("error waiting on container %q: %v", cid, err)
	}
	return ws, nil
}

// WaitPID waits for process 'pid' in the container's sandbox and returns its
// WaitStatus.
func (s *Sandbox) WaitPID(pid int32, cid string) (syscall.WaitStatus, error) {
	log.Debugf("Waiting for PID %d in sandbox %q", pid, s.ID)
	var ws syscall.WaitStatus
	conn, err := s.connect()
	if err != nil {
		return ws, err
	}
	defer conn.Close()

	args := &boot.WaitPIDArgs{
		PID: pid,
		CID: cid,
	}

	if err := conn.Call(boot.ContainerWaitPID, args, &ws); err != nil {
		return ws, fmt.Errorf("error waiting on PID %d in sandbox %q: %v", pid, s.ID, err)
	}

	if s.IsRootContainer(cid) {
		// If waiting for the root, give some time for the sandbox process to exit
		// to prevent races with resources that might still be in use.
		timeout := time.Now().Add(time.Second)
		log.Debugf("Waiting for the sandbox process to exit")
		for s.IsRunning() {
			if time.Now().After(timeout) {
				log.Debugf("Timeout waiting for sandbox process to exit")
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
	return ws, nil
}

// IsRootContainer returns true if the specified container ID belongs to the
// root container.
func (s *Sandbox) IsRootContainer(cid string) bool {
	return s.ID == cid
}

// Stop stops the container in the sandbox.
func (s *Sandbox) Stop(cid string) error {
	// TODO: This should stop the container with the given ID
	// in the sandbox.
	return nil
}

// Destroy frees all resources associated with the sandbox.
func (s *Sandbox) Destroy() error {
	log.Debugf("Destroy sandbox %q", s.ID)
	if s.Pid != 0 {
		// TODO: Too harsh?
		log.Debugf("Killing sandbox %q", s.ID)
		killProcess(s.Pid, unix.SIGKILL)
		s.Pid = 0
	}
	if s.GoferPid != 0 {
		log.Debugf("Killing gofer for sandbox %q", s.ID)
		killProcess(s.GoferPid, unix.SIGKILL)
		s.GoferPid = 0
	}

	return nil
}

// Signal sends the signal to a container in the sandbox.
func (s *Sandbox) Signal(cid string, sig syscall.Signal) error {
	log.Debugf("Signal sandbox %q", s.ID)
	conn, err := s.connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	args := boot.SignalArgs{
		CID:   cid,
		Signo: int32(sig),
	}
	if err := conn.Call(boot.ContainerSignal, &args, nil); err != nil {
		return fmt.Errorf("err signaling container %q: %v", cid, err)
	}
	return nil
}

// Checkpoint sends the checkpoint call for a container in the sandbox.
// The statefile will be written to f.
func (s *Sandbox) Checkpoint(cid string, f *os.File) error {
	log.Debugf("Checkpoint sandbox %q", s.ID)
	conn, err := s.connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	opt := control.SaveOpts{
		FilePayload: urpc.FilePayload{
			Files: []*os.File{f},
		},
	}

	if err := conn.Call(boot.ContainerCheckpoint, &opt, nil); err != nil {
		return fmt.Errorf("err checkpointing container %q: %v", cid, err)
	}
	return nil
}

// Pause sends the pause call for a container in the sandbox.
func (s *Sandbox) Pause(cid string) error {
	log.Debugf("Pause sandbox %q", s.ID)
	conn, err := s.connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	if err := conn.Call(boot.ContainerPause, nil, nil); err != nil {
		return fmt.Errorf("err pausing container %q: %v", cid, err)
	}
	return nil
}

// Resume sends the resume call for a container in the sandbox.
func (s *Sandbox) Resume(cid string) error {
	log.Debugf("Resume sandbox %q", s.ID)
	conn, err := s.connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	if err := conn.Call(boot.ContainerResume, nil, nil); err != nil {
		return fmt.Errorf("err resuming container %q: %v", cid, err)
	}
	return nil
}

// IsRunning returns true if the sandbox or gofer process is running.
func (s *Sandbox) IsRunning() bool {
	if s.Pid != 0 {
		// Send a signal 0 to the sandbox process.
		if err := killProcess(s.Pid, 0); err == nil {
			return true
		}
	}
	if s.GoferPid != 0 {
		// Send a signal 0 to the gofer process.
		if err := killProcess(s.GoferPid, 0); err == nil {
			log.Warningf("Found orphan gofer process, pid: %d", s.GoferPid)
			// Attempt to kill gofer if it's orphan.
			killProcess(s.GoferPid, unix.SIGKILL)

			// Don't wait for gofer to die. Return 'running' and hope gofer is dead
			// next time around.
			return true
		}
	}
	return false
}

// Stacks collects and returns all stacks for the sandbox.
func (s *Sandbox) Stacks() (string, error) {
	log.Debugf("Stacks sandbox %q", s.ID)
	conn, err := s.connect()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	var stacks string
	if err := conn.Call(boot.SandboxStacks, nil, &stacks); err != nil {
		return "", fmt.Errorf("err getting sandbox %q stacks: %v", s.ID, err)
	}
	return stacks, nil
}

// killProcess sends a signal to the host process (i.e. a sandbox or gofer
// process). Sandbox.Signal should be used to send a signal to a process
// running inside the sandbox.
func killProcess(pid int, sig syscall.Signal) error {
	if err := syscall.Kill(pid, sig); err != nil {
		return fmt.Errorf("error sending signal %d to pid %d: %v", sig, pid, err)
	}
	return nil
}
