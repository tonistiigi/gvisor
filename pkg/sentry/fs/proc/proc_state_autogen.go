// automatically generated by stateify.

package proc

import (
	"gvisor.googlesource.com/gvisor/pkg/state"
)

func (x *execArgInode) beforeSave() {}
func (x *execArgInode) save(m state.Map) {
	x.beforeSave()
	m.Save("SimpleFileInode", &x.SimpleFileInode)
	m.Save("arg", &x.arg)
	m.Save("t", &x.t)
}

func (x *execArgInode) afterLoad() {}
func (x *execArgInode) load(m state.Map) {
	m.Load("SimpleFileInode", &x.SimpleFileInode)
	m.Load("arg", &x.arg)
	m.Load("t", &x.t)
}

func (x *execArgFile) beforeSave() {}
func (x *execArgFile) save(m state.Map) {
	x.beforeSave()
	m.Save("arg", &x.arg)
	m.Save("t", &x.t)
}

func (x *execArgFile) afterLoad() {}
func (x *execArgFile) load(m state.Map) {
	m.Load("arg", &x.arg)
	m.Load("t", &x.t)
}

func (x *fdDir) beforeSave() {}
func (x *fdDir) save(m state.Map) {
	x.beforeSave()
	m.Save("Dir", &x.Dir)
	m.Save("t", &x.t)
}

func (x *fdDir) afterLoad() {}
func (x *fdDir) load(m state.Map) {
	m.Load("Dir", &x.Dir)
	m.Load("t", &x.t)
}

func (x *fdDirFile) beforeSave() {}
func (x *fdDirFile) save(m state.Map) {
	x.beforeSave()
	m.Save("isInfoFile", &x.isInfoFile)
	m.Save("t", &x.t)
}

func (x *fdDirFile) afterLoad() {}
func (x *fdDirFile) load(m state.Map) {
	m.Load("isInfoFile", &x.isInfoFile)
	m.Load("t", &x.t)
}

func (x *fdInfoDir) beforeSave() {}
func (x *fdInfoDir) save(m state.Map) {
	x.beforeSave()
	m.Save("Dir", &x.Dir)
	m.Save("t", &x.t)
}

func (x *fdInfoDir) afterLoad() {}
func (x *fdInfoDir) load(m state.Map) {
	m.Load("Dir", &x.Dir)
	m.Load("t", &x.t)
}

func (x *filesystemsData) beforeSave() {}
func (x *filesystemsData) save(m state.Map) {
	x.beforeSave()
}

func (x *filesystemsData) afterLoad() {}
func (x *filesystemsData) load(m state.Map) {
}

func (x *filesystem) beforeSave() {}
func (x *filesystem) save(m state.Map) {
	x.beforeSave()
}

func (x *filesystem) afterLoad() {}
func (x *filesystem) load(m state.Map) {
}

func (x *taskOwnedInodeOps) beforeSave() {}
func (x *taskOwnedInodeOps) save(m state.Map) {
	x.beforeSave()
	m.Save("InodeOperations", &x.InodeOperations)
	m.Save("t", &x.t)
}

func (x *taskOwnedInodeOps) afterLoad() {}
func (x *taskOwnedInodeOps) load(m state.Map) {
	m.Load("InodeOperations", &x.InodeOperations)
	m.Load("t", &x.t)
}

func (x *staticFileInodeOps) beforeSave() {}
func (x *staticFileInodeOps) save(m state.Map) {
	x.beforeSave()
	m.Save("InodeSimpleAttributes", &x.InodeSimpleAttributes)
	m.Save("InodeStaticFileGetter", &x.InodeStaticFileGetter)
}

func (x *staticFileInodeOps) afterLoad() {}
func (x *staticFileInodeOps) load(m state.Map) {
	m.Load("InodeSimpleAttributes", &x.InodeSimpleAttributes)
	m.Load("InodeStaticFileGetter", &x.InodeStaticFileGetter)
}

func (x *loadavgData) beforeSave() {}
func (x *loadavgData) save(m state.Map) {
	x.beforeSave()
}

func (x *loadavgData) afterLoad() {}
func (x *loadavgData) load(m state.Map) {
}

func (x *meminfoData) beforeSave() {}
func (x *meminfoData) save(m state.Map) {
	x.beforeSave()
	m.Save("k", &x.k)
}

func (x *meminfoData) afterLoad() {}
func (x *meminfoData) load(m state.Map) {
	m.Load("k", &x.k)
}

func (x *mountInfoFile) beforeSave() {}
func (x *mountInfoFile) save(m state.Map) {
	x.beforeSave()
	m.Save("t", &x.t)
}

func (x *mountInfoFile) afterLoad() {}
func (x *mountInfoFile) load(m state.Map) {
	m.Load("t", &x.t)
}

func (x *mountsFile) beforeSave() {}
func (x *mountsFile) save(m state.Map) {
	x.beforeSave()
	m.Save("t", &x.t)
}

func (x *mountsFile) afterLoad() {}
func (x *mountsFile) load(m state.Map) {
	m.Load("t", &x.t)
}

func (x *ifinet6) beforeSave() {}
func (x *ifinet6) save(m state.Map) {
	x.beforeSave()
	m.Save("s", &x.s)
}

func (x *ifinet6) afterLoad() {}
func (x *ifinet6) load(m state.Map) {
	m.Load("s", &x.s)
}

func (x *netDev) beforeSave() {}
func (x *netDev) save(m state.Map) {
	x.beforeSave()
	m.Save("s", &x.s)
}

func (x *netDev) afterLoad() {}
func (x *netDev) load(m state.Map) {
	m.Load("s", &x.s)
}

func (x *netUnix) beforeSave() {}
func (x *netUnix) save(m state.Map) {
	x.beforeSave()
	m.Save("k", &x.k)
}

func (x *netUnix) afterLoad() {}
func (x *netUnix) load(m state.Map) {
	m.Load("k", &x.k)
}

func (x *proc) beforeSave() {}
func (x *proc) save(m state.Map) {
	x.beforeSave()
	m.Save("Dir", &x.Dir)
	m.Save("k", &x.k)
	m.Save("pidns", &x.pidns)
}

func (x *proc) afterLoad() {}
func (x *proc) load(m state.Map) {
	m.Load("Dir", &x.Dir)
	m.Load("k", &x.k)
	m.Load("pidns", &x.pidns)
}

func (x *self) beforeSave() {}
func (x *self) save(m state.Map) {
	x.beforeSave()
	m.Save("Symlink", &x.Symlink)
	m.Save("pidns", &x.pidns)
}

func (x *self) afterLoad() {}
func (x *self) load(m state.Map) {
	m.Load("Symlink", &x.Symlink)
	m.Load("pidns", &x.pidns)
}

func (x *threadSelf) beforeSave() {}
func (x *threadSelf) save(m state.Map) {
	x.beforeSave()
	m.Save("Symlink", &x.Symlink)
	m.Save("pidns", &x.pidns)
}

func (x *threadSelf) afterLoad() {}
func (x *threadSelf) load(m state.Map) {
	m.Load("Symlink", &x.Symlink)
	m.Load("pidns", &x.pidns)
}

func (x *rootProcFile) beforeSave() {}
func (x *rootProcFile) save(m state.Map) {
	x.beforeSave()
	m.Save("iops", &x.iops)
}

func (x *rootProcFile) afterLoad() {}
func (x *rootProcFile) load(m state.Map) {
	m.Load("iops", &x.iops)
}

func (x *statData) beforeSave() {}
func (x *statData) save(m state.Map) {
	x.beforeSave()
	m.Save("k", &x.k)
}

func (x *statData) afterLoad() {}
func (x *statData) load(m state.Map) {
	m.Load("k", &x.k)
}

func (x *mmapMinAddrData) beforeSave() {}
func (x *mmapMinAddrData) save(m state.Map) {
	x.beforeSave()
	m.Save("k", &x.k)
}

func (x *mmapMinAddrData) afterLoad() {}
func (x *mmapMinAddrData) load(m state.Map) {
	m.Load("k", &x.k)
}

func (x *overcommitMemory) beforeSave() {}
func (x *overcommitMemory) save(m state.Map) {
	x.beforeSave()
}

func (x *overcommitMemory) afterLoad() {}
func (x *overcommitMemory) load(m state.Map) {
}

func (x *hostname) beforeSave() {}
func (x *hostname) save(m state.Map) {
	x.beforeSave()
	m.Save("SimpleFileInode", &x.SimpleFileInode)
}

func (x *hostname) afterLoad() {}
func (x *hostname) load(m state.Map) {
	m.Load("SimpleFileInode", &x.SimpleFileInode)
}

func (x *hostnameFile) beforeSave() {}
func (x *hostnameFile) save(m state.Map) {
	x.beforeSave()
}

func (x *hostnameFile) afterLoad() {}
func (x *hostnameFile) load(m state.Map) {
}

func (x *tcpMemInode) save(m state.Map) {
	x.beforeSave()
	m.Save("SimpleFileInode", &x.SimpleFileInode)
	m.Save("dir", &x.dir)
	m.Save("s", &x.s)
	m.Save("size", &x.size)
}

func (x *tcpMemInode) load(m state.Map) {
	m.Load("SimpleFileInode", &x.SimpleFileInode)
	m.Load("dir", &x.dir)
	m.LoadWait("s", &x.s)
	m.Load("size", &x.size)
	m.AfterLoad(x.afterLoad)
}

func (x *tcpMemFile) beforeSave() {}
func (x *tcpMemFile) save(m state.Map) {
	x.beforeSave()
	m.Save("tcpMemInode", &x.tcpMemInode)
}

func (x *tcpMemFile) afterLoad() {}
func (x *tcpMemFile) load(m state.Map) {
	m.Load("tcpMemInode", &x.tcpMemInode)
}

func (x *tcpSack) beforeSave() {}
func (x *tcpSack) save(m state.Map) {
	x.beforeSave()
	m.Save("stack", &x.stack)
	m.Save("enabled", &x.enabled)
	m.Save("SimpleFileInode", &x.SimpleFileInode)
}

func (x *tcpSack) load(m state.Map) {
	m.LoadWait("stack", &x.stack)
	m.Load("enabled", &x.enabled)
	m.Load("SimpleFileInode", &x.SimpleFileInode)
	m.AfterLoad(x.afterLoad)
}

func (x *tcpSackFile) beforeSave() {}
func (x *tcpSackFile) save(m state.Map) {
	x.beforeSave()
	m.Save("tcpSack", &x.tcpSack)
	m.Save("stack", &x.stack)
}

func (x *tcpSackFile) afterLoad() {}
func (x *tcpSackFile) load(m state.Map) {
	m.Load("tcpSack", &x.tcpSack)
	m.LoadWait("stack", &x.stack)
}

func (x *taskDir) beforeSave() {}
func (x *taskDir) save(m state.Map) {
	x.beforeSave()
	m.Save("Dir", &x.Dir)
	m.Save("t", &x.t)
	m.Save("pidns", &x.pidns)
}

func (x *taskDir) afterLoad() {}
func (x *taskDir) load(m state.Map) {
	m.Load("Dir", &x.Dir)
	m.Load("t", &x.t)
	m.Load("pidns", &x.pidns)
}

func (x *subtasks) beforeSave() {}
func (x *subtasks) save(m state.Map) {
	x.beforeSave()
	m.Save("Dir", &x.Dir)
	m.Save("t", &x.t)
	m.Save("pidns", &x.pidns)
}

func (x *subtasks) afterLoad() {}
func (x *subtasks) load(m state.Map) {
	m.Load("Dir", &x.Dir)
	m.Load("t", &x.t)
	m.Load("pidns", &x.pidns)
}

func (x *subtasksFile) beforeSave() {}
func (x *subtasksFile) save(m state.Map) {
	x.beforeSave()
	m.Save("t", &x.t)
	m.Save("pidns", &x.pidns)
}

func (x *subtasksFile) afterLoad() {}
func (x *subtasksFile) load(m state.Map) {
	m.Load("t", &x.t)
	m.Load("pidns", &x.pidns)
}

func (x *exe) beforeSave() {}
func (x *exe) save(m state.Map) {
	x.beforeSave()
	m.Save("Symlink", &x.Symlink)
	m.Save("t", &x.t)
}

func (x *exe) afterLoad() {}
func (x *exe) load(m state.Map) {
	m.Load("Symlink", &x.Symlink)
	m.Load("t", &x.t)
}

func (x *namespaceSymlink) beforeSave() {}
func (x *namespaceSymlink) save(m state.Map) {
	x.beforeSave()
	m.Save("Symlink", &x.Symlink)
	m.Save("t", &x.t)
}

func (x *namespaceSymlink) afterLoad() {}
func (x *namespaceSymlink) load(m state.Map) {
	m.Load("Symlink", &x.Symlink)
	m.Load("t", &x.t)
}

func (x *mapsData) beforeSave() {}
func (x *mapsData) save(m state.Map) {
	x.beforeSave()
	m.Save("t", &x.t)
}

func (x *mapsData) afterLoad() {}
func (x *mapsData) load(m state.Map) {
	m.Load("t", &x.t)
}

func (x *smapsData) beforeSave() {}
func (x *smapsData) save(m state.Map) {
	x.beforeSave()
	m.Save("t", &x.t)
}

func (x *smapsData) afterLoad() {}
func (x *smapsData) load(m state.Map) {
	m.Load("t", &x.t)
}

func (x *taskStatData) beforeSave() {}
func (x *taskStatData) save(m state.Map) {
	x.beforeSave()
	m.Save("t", &x.t)
	m.Save("tgstats", &x.tgstats)
	m.Save("pidns", &x.pidns)
}

func (x *taskStatData) afterLoad() {}
func (x *taskStatData) load(m state.Map) {
	m.Load("t", &x.t)
	m.Load("tgstats", &x.tgstats)
	m.Load("pidns", &x.pidns)
}

func (x *statmData) beforeSave() {}
func (x *statmData) save(m state.Map) {
	x.beforeSave()
	m.Save("t", &x.t)
}

func (x *statmData) afterLoad() {}
func (x *statmData) load(m state.Map) {
	m.Load("t", &x.t)
}

func (x *statusData) beforeSave() {}
func (x *statusData) save(m state.Map) {
	x.beforeSave()
	m.Save("t", &x.t)
	m.Save("pidns", &x.pidns)
}

func (x *statusData) afterLoad() {}
func (x *statusData) load(m state.Map) {
	m.Load("t", &x.t)
	m.Load("pidns", &x.pidns)
}

func (x *ioData) beforeSave() {}
func (x *ioData) save(m state.Map) {
	x.beforeSave()
	m.Save("ioUsage", &x.ioUsage)
}

func (x *ioData) afterLoad() {}
func (x *ioData) load(m state.Map) {
	m.Load("ioUsage", &x.ioUsage)
}

func (x *comm) beforeSave() {}
func (x *comm) save(m state.Map) {
	x.beforeSave()
	m.Save("SimpleFileInode", &x.SimpleFileInode)
	m.Save("t", &x.t)
}

func (x *comm) afterLoad() {}
func (x *comm) load(m state.Map) {
	m.Load("SimpleFileInode", &x.SimpleFileInode)
	m.Load("t", &x.t)
}

func (x *commFile) beforeSave() {}
func (x *commFile) save(m state.Map) {
	x.beforeSave()
	m.Save("t", &x.t)
}

func (x *commFile) afterLoad() {}
func (x *commFile) load(m state.Map) {
	m.Load("t", &x.t)
}

func (x *auxvec) beforeSave() {}
func (x *auxvec) save(m state.Map) {
	x.beforeSave()
	m.Save("SimpleFileInode", &x.SimpleFileInode)
	m.Save("t", &x.t)
}

func (x *auxvec) afterLoad() {}
func (x *auxvec) load(m state.Map) {
	m.Load("SimpleFileInode", &x.SimpleFileInode)
	m.Load("t", &x.t)
}

func (x *auxvecFile) beforeSave() {}
func (x *auxvecFile) save(m state.Map) {
	x.beforeSave()
	m.Save("t", &x.t)
}

func (x *auxvecFile) afterLoad() {}
func (x *auxvecFile) load(m state.Map) {
	m.Load("t", &x.t)
}

func (x *idMapInodeOperations) beforeSave() {}
func (x *idMapInodeOperations) save(m state.Map) {
	x.beforeSave()
	m.Save("InodeSimpleAttributes", &x.InodeSimpleAttributes)
	m.Save("InodeSimpleExtendedAttributes", &x.InodeSimpleExtendedAttributes)
	m.Save("t", &x.t)
	m.Save("gids", &x.gids)
}

func (x *idMapInodeOperations) afterLoad() {}
func (x *idMapInodeOperations) load(m state.Map) {
	m.Load("InodeSimpleAttributes", &x.InodeSimpleAttributes)
	m.Load("InodeSimpleExtendedAttributes", &x.InodeSimpleExtendedAttributes)
	m.Load("t", &x.t)
	m.Load("gids", &x.gids)
}

func (x *idMapFileOperations) beforeSave() {}
func (x *idMapFileOperations) save(m state.Map) {
	x.beforeSave()
	m.Save("iops", &x.iops)
}

func (x *idMapFileOperations) afterLoad() {}
func (x *idMapFileOperations) load(m state.Map) {
	m.Load("iops", &x.iops)
}

func (x *uptime) beforeSave() {}
func (x *uptime) save(m state.Map) {
	x.beforeSave()
	m.Save("SimpleFileInode", &x.SimpleFileInode)
	m.Save("startTime", &x.startTime)
}

func (x *uptime) afterLoad() {}
func (x *uptime) load(m state.Map) {
	m.Load("SimpleFileInode", &x.SimpleFileInode)
	m.Load("startTime", &x.startTime)
}

func (x *uptimeFile) beforeSave() {}
func (x *uptimeFile) save(m state.Map) {
	x.beforeSave()
	m.Save("startTime", &x.startTime)
}

func (x *uptimeFile) afterLoad() {}
func (x *uptimeFile) load(m state.Map) {
	m.Load("startTime", &x.startTime)
}

func (x *versionData) beforeSave() {}
func (x *versionData) save(m state.Map) {
	x.beforeSave()
	m.Save("k", &x.k)
}

func (x *versionData) afterLoad() {}
func (x *versionData) load(m state.Map) {
	m.Load("k", &x.k)
}

func init() {
	state.Register("proc.execArgInode", (*execArgInode)(nil), state.Fns{Save: (*execArgInode).save, Load: (*execArgInode).load})
	state.Register("proc.execArgFile", (*execArgFile)(nil), state.Fns{Save: (*execArgFile).save, Load: (*execArgFile).load})
	state.Register("proc.fdDir", (*fdDir)(nil), state.Fns{Save: (*fdDir).save, Load: (*fdDir).load})
	state.Register("proc.fdDirFile", (*fdDirFile)(nil), state.Fns{Save: (*fdDirFile).save, Load: (*fdDirFile).load})
	state.Register("proc.fdInfoDir", (*fdInfoDir)(nil), state.Fns{Save: (*fdInfoDir).save, Load: (*fdInfoDir).load})
	state.Register("proc.filesystemsData", (*filesystemsData)(nil), state.Fns{Save: (*filesystemsData).save, Load: (*filesystemsData).load})
	state.Register("proc.filesystem", (*filesystem)(nil), state.Fns{Save: (*filesystem).save, Load: (*filesystem).load})
	state.Register("proc.taskOwnedInodeOps", (*taskOwnedInodeOps)(nil), state.Fns{Save: (*taskOwnedInodeOps).save, Load: (*taskOwnedInodeOps).load})
	state.Register("proc.staticFileInodeOps", (*staticFileInodeOps)(nil), state.Fns{Save: (*staticFileInodeOps).save, Load: (*staticFileInodeOps).load})
	state.Register("proc.loadavgData", (*loadavgData)(nil), state.Fns{Save: (*loadavgData).save, Load: (*loadavgData).load})
	state.Register("proc.meminfoData", (*meminfoData)(nil), state.Fns{Save: (*meminfoData).save, Load: (*meminfoData).load})
	state.Register("proc.mountInfoFile", (*mountInfoFile)(nil), state.Fns{Save: (*mountInfoFile).save, Load: (*mountInfoFile).load})
	state.Register("proc.mountsFile", (*mountsFile)(nil), state.Fns{Save: (*mountsFile).save, Load: (*mountsFile).load})
	state.Register("proc.ifinet6", (*ifinet6)(nil), state.Fns{Save: (*ifinet6).save, Load: (*ifinet6).load})
	state.Register("proc.netDev", (*netDev)(nil), state.Fns{Save: (*netDev).save, Load: (*netDev).load})
	state.Register("proc.netUnix", (*netUnix)(nil), state.Fns{Save: (*netUnix).save, Load: (*netUnix).load})
	state.Register("proc.proc", (*proc)(nil), state.Fns{Save: (*proc).save, Load: (*proc).load})
	state.Register("proc.self", (*self)(nil), state.Fns{Save: (*self).save, Load: (*self).load})
	state.Register("proc.threadSelf", (*threadSelf)(nil), state.Fns{Save: (*threadSelf).save, Load: (*threadSelf).load})
	state.Register("proc.rootProcFile", (*rootProcFile)(nil), state.Fns{Save: (*rootProcFile).save, Load: (*rootProcFile).load})
	state.Register("proc.statData", (*statData)(nil), state.Fns{Save: (*statData).save, Load: (*statData).load})
	state.Register("proc.mmapMinAddrData", (*mmapMinAddrData)(nil), state.Fns{Save: (*mmapMinAddrData).save, Load: (*mmapMinAddrData).load})
	state.Register("proc.overcommitMemory", (*overcommitMemory)(nil), state.Fns{Save: (*overcommitMemory).save, Load: (*overcommitMemory).load})
	state.Register("proc.hostname", (*hostname)(nil), state.Fns{Save: (*hostname).save, Load: (*hostname).load})
	state.Register("proc.hostnameFile", (*hostnameFile)(nil), state.Fns{Save: (*hostnameFile).save, Load: (*hostnameFile).load})
	state.Register("proc.tcpMemInode", (*tcpMemInode)(nil), state.Fns{Save: (*tcpMemInode).save, Load: (*tcpMemInode).load})
	state.Register("proc.tcpMemFile", (*tcpMemFile)(nil), state.Fns{Save: (*tcpMemFile).save, Load: (*tcpMemFile).load})
	state.Register("proc.tcpSack", (*tcpSack)(nil), state.Fns{Save: (*tcpSack).save, Load: (*tcpSack).load})
	state.Register("proc.tcpSackFile", (*tcpSackFile)(nil), state.Fns{Save: (*tcpSackFile).save, Load: (*tcpSackFile).load})
	state.Register("proc.taskDir", (*taskDir)(nil), state.Fns{Save: (*taskDir).save, Load: (*taskDir).load})
	state.Register("proc.subtasks", (*subtasks)(nil), state.Fns{Save: (*subtasks).save, Load: (*subtasks).load})
	state.Register("proc.subtasksFile", (*subtasksFile)(nil), state.Fns{Save: (*subtasksFile).save, Load: (*subtasksFile).load})
	state.Register("proc.exe", (*exe)(nil), state.Fns{Save: (*exe).save, Load: (*exe).load})
	state.Register("proc.namespaceSymlink", (*namespaceSymlink)(nil), state.Fns{Save: (*namespaceSymlink).save, Load: (*namespaceSymlink).load})
	state.Register("proc.mapsData", (*mapsData)(nil), state.Fns{Save: (*mapsData).save, Load: (*mapsData).load})
	state.Register("proc.smapsData", (*smapsData)(nil), state.Fns{Save: (*smapsData).save, Load: (*smapsData).load})
	state.Register("proc.taskStatData", (*taskStatData)(nil), state.Fns{Save: (*taskStatData).save, Load: (*taskStatData).load})
	state.Register("proc.statmData", (*statmData)(nil), state.Fns{Save: (*statmData).save, Load: (*statmData).load})
	state.Register("proc.statusData", (*statusData)(nil), state.Fns{Save: (*statusData).save, Load: (*statusData).load})
	state.Register("proc.ioData", (*ioData)(nil), state.Fns{Save: (*ioData).save, Load: (*ioData).load})
	state.Register("proc.comm", (*comm)(nil), state.Fns{Save: (*comm).save, Load: (*comm).load})
	state.Register("proc.commFile", (*commFile)(nil), state.Fns{Save: (*commFile).save, Load: (*commFile).load})
	state.Register("proc.auxvec", (*auxvec)(nil), state.Fns{Save: (*auxvec).save, Load: (*auxvec).load})
	state.Register("proc.auxvecFile", (*auxvecFile)(nil), state.Fns{Save: (*auxvecFile).save, Load: (*auxvecFile).load})
	state.Register("proc.idMapInodeOperations", (*idMapInodeOperations)(nil), state.Fns{Save: (*idMapInodeOperations).save, Load: (*idMapInodeOperations).load})
	state.Register("proc.idMapFileOperations", (*idMapFileOperations)(nil), state.Fns{Save: (*idMapFileOperations).save, Load: (*idMapFileOperations).load})
	state.Register("proc.uptime", (*uptime)(nil), state.Fns{Save: (*uptime).save, Load: (*uptime).load})
	state.Register("proc.uptimeFile", (*uptimeFile)(nil), state.Fns{Save: (*uptimeFile).save, Load: (*uptimeFile).load})
	state.Register("proc.versionData", (*versionData)(nil), state.Fns{Save: (*versionData).save, Load: (*versionData).load})
}
