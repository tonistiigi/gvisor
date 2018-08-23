// automatically generated by stateify.

package fsutil

import (
	"gvisor.googlesource.com/gvisor/pkg/state"
)

func (x *DirtySet) beforeSave() {}
func (x *DirtySet) save(m state.Map) {
	x.beforeSave()
	var root *DirtySegmentDataSlices = x.saveRoot()
	m.SaveValue("root", root)
}

func (x *DirtySet) afterLoad() {}
func (x *DirtySet) load(m state.Map) {
	m.LoadValue("root", new(*DirtySegmentDataSlices), func(y interface{}) { x.loadRoot(y.(*DirtySegmentDataSlices)) })
}

func (x *Dirtynode) beforeSave() {}
func (x *Dirtynode) save(m state.Map) {
	x.beforeSave()
	m.Save("nrSegments", &x.nrSegments)
	m.Save("parent", &x.parent)
	m.Save("parentIndex", &x.parentIndex)
	m.Save("hasChildren", &x.hasChildren)
	m.Save("keys", &x.keys)
	m.Save("values", &x.values)
	m.Save("children", &x.children)
}

func (x *Dirtynode) afterLoad() {}
func (x *Dirtynode) load(m state.Map) {
	m.Load("nrSegments", &x.nrSegments)
	m.Load("parent", &x.parent)
	m.Load("parentIndex", &x.parentIndex)
	m.Load("hasChildren", &x.hasChildren)
	m.Load("keys", &x.keys)
	m.Load("values", &x.values)
	m.Load("children", &x.children)
}

func (x *DirtyIterator) beforeSave() {}
func (x *DirtyIterator) save(m state.Map) {
	x.beforeSave()
	m.Save("node", &x.node)
	m.Save("index", &x.index)
}

func (x *DirtyIterator) afterLoad() {}
func (x *DirtyIterator) load(m state.Map) {
	m.Load("node", &x.node)
	m.Load("index", &x.index)
}

func (x *DirtyGapIterator) beforeSave() {}
func (x *DirtyGapIterator) save(m state.Map) {
	x.beforeSave()
	m.Save("node", &x.node)
	m.Save("index", &x.index)
}

func (x *DirtyGapIterator) afterLoad() {}
func (x *DirtyGapIterator) load(m state.Map) {
	m.Load("node", &x.node)
	m.Load("index", &x.index)
}

func (x *DirtySegmentDataSlices) beforeSave() {}
func (x *DirtySegmentDataSlices) save(m state.Map) {
	x.beforeSave()
	m.Save("Start", &x.Start)
	m.Save("End", &x.End)
	m.Save("Values", &x.Values)
}

func (x *DirtySegmentDataSlices) afterLoad() {}
func (x *DirtySegmentDataSlices) load(m state.Map) {
	m.Load("Start", &x.Start)
	m.Load("End", &x.End)
	m.Load("Values", &x.Values)
}

func (x *NoopRelease) beforeSave() {}
func (x *NoopRelease) save(m state.Map) {
	x.beforeSave()
}

func (x *NoopRelease) afterLoad() {}
func (x *NoopRelease) load(m state.Map) {
}

func (x *GenericSeek) beforeSave() {}
func (x *GenericSeek) save(m state.Map) {
	x.beforeSave()
}

func (x *GenericSeek) afterLoad() {}
func (x *GenericSeek) load(m state.Map) {
}

func (x *ZeroSeek) beforeSave() {}
func (x *ZeroSeek) save(m state.Map) {
	x.beforeSave()
}

func (x *ZeroSeek) afterLoad() {}
func (x *ZeroSeek) load(m state.Map) {
}

func (x *PipeSeek) beforeSave() {}
func (x *PipeSeek) save(m state.Map) {
	x.beforeSave()
}

func (x *PipeSeek) afterLoad() {}
func (x *PipeSeek) load(m state.Map) {
}

func (x *NotDirReaddir) beforeSave() {}
func (x *NotDirReaddir) save(m state.Map) {
	x.beforeSave()
}

func (x *NotDirReaddir) afterLoad() {}
func (x *NotDirReaddir) load(m state.Map) {
}

func (x *NoFsync) beforeSave() {}
func (x *NoFsync) save(m state.Map) {
	x.beforeSave()
}

func (x *NoFsync) afterLoad() {}
func (x *NoFsync) load(m state.Map) {
}

func (x *NoopFsync) beforeSave() {}
func (x *NoopFsync) save(m state.Map) {
	x.beforeSave()
}

func (x *NoopFsync) afterLoad() {}
func (x *NoopFsync) load(m state.Map) {
}

func (x *NoopFlush) beforeSave() {}
func (x *NoopFlush) save(m state.Map) {
	x.beforeSave()
}

func (x *NoopFlush) afterLoad() {}
func (x *NoopFlush) load(m state.Map) {
}

func (x *NoMMap) beforeSave() {}
func (x *NoMMap) save(m state.Map) {
	x.beforeSave()
}

func (x *NoMMap) afterLoad() {}
func (x *NoMMap) load(m state.Map) {
}

func (x *NoIoctl) beforeSave() {}
func (x *NoIoctl) save(m state.Map) {
	x.beforeSave()
}

func (x *NoIoctl) afterLoad() {}
func (x *NoIoctl) load(m state.Map) {
}

func (x *DirFileOperations) beforeSave() {}
func (x *DirFileOperations) save(m state.Map) {
	x.beforeSave()
	m.Save("dentryMap", &x.dentryMap)
	m.Save("dirCursor", &x.dirCursor)
}

func (x *DirFileOperations) afterLoad() {}
func (x *DirFileOperations) load(m state.Map) {
	m.Load("dentryMap", &x.dentryMap)
	m.Load("dirCursor", &x.dirCursor)
}

func (x *FileRangeSet) beforeSave() {}
func (x *FileRangeSet) save(m state.Map) {
	x.beforeSave()
	var root *FileRangeSegmentDataSlices = x.saveRoot()
	m.SaveValue("root", root)
}

func (x *FileRangeSet) afterLoad() {}
func (x *FileRangeSet) load(m state.Map) {
	m.LoadValue("root", new(*FileRangeSegmentDataSlices), func(y interface{}) { x.loadRoot(y.(*FileRangeSegmentDataSlices)) })
}

func (x *FileRangenode) beforeSave() {}
func (x *FileRangenode) save(m state.Map) {
	x.beforeSave()
	m.Save("nrSegments", &x.nrSegments)
	m.Save("parent", &x.parent)
	m.Save("parentIndex", &x.parentIndex)
	m.Save("hasChildren", &x.hasChildren)
	m.Save("keys", &x.keys)
	m.Save("values", &x.values)
	m.Save("children", &x.children)
}

func (x *FileRangenode) afterLoad() {}
func (x *FileRangenode) load(m state.Map) {
	m.Load("nrSegments", &x.nrSegments)
	m.Load("parent", &x.parent)
	m.Load("parentIndex", &x.parentIndex)
	m.Load("hasChildren", &x.hasChildren)
	m.Load("keys", &x.keys)
	m.Load("values", &x.values)
	m.Load("children", &x.children)
}

func (x *FileRangeIterator) beforeSave() {}
func (x *FileRangeIterator) save(m state.Map) {
	x.beforeSave()
	m.Save("node", &x.node)
	m.Save("index", &x.index)
}

func (x *FileRangeIterator) afterLoad() {}
func (x *FileRangeIterator) load(m state.Map) {
	m.Load("node", &x.node)
	m.Load("index", &x.index)
}

func (x *FileRangeGapIterator) beforeSave() {}
func (x *FileRangeGapIterator) save(m state.Map) {
	x.beforeSave()
	m.Save("node", &x.node)
	m.Save("index", &x.index)
}

func (x *FileRangeGapIterator) afterLoad() {}
func (x *FileRangeGapIterator) load(m state.Map) {
	m.Load("node", &x.node)
	m.Load("index", &x.index)
}

func (x *FileRangeSegmentDataSlices) beforeSave() {}
func (x *FileRangeSegmentDataSlices) save(m state.Map) {
	x.beforeSave()
	m.Save("Start", &x.Start)
	m.Save("End", &x.End)
	m.Save("Values", &x.Values)
}

func (x *FileRangeSegmentDataSlices) afterLoad() {}
func (x *FileRangeSegmentDataSlices) load(m state.Map) {
	m.Load("Start", &x.Start)
	m.Load("End", &x.End)
	m.Load("Values", &x.Values)
}

func (x *frameRefSet) beforeSave() {}
func (x *frameRefSet) save(m state.Map) {
	x.beforeSave()
	var root *frameRefSegmentDataSlices = x.saveRoot()
	m.SaveValue("root", root)
}

func (x *frameRefSet) afterLoad() {}
func (x *frameRefSet) load(m state.Map) {
	m.LoadValue("root", new(*frameRefSegmentDataSlices), func(y interface{}) { x.loadRoot(y.(*frameRefSegmentDataSlices)) })
}

func (x *frameRefnode) beforeSave() {}
func (x *frameRefnode) save(m state.Map) {
	x.beforeSave()
	m.Save("nrSegments", &x.nrSegments)
	m.Save("parent", &x.parent)
	m.Save("parentIndex", &x.parentIndex)
	m.Save("hasChildren", &x.hasChildren)
	m.Save("keys", &x.keys)
	m.Save("values", &x.values)
	m.Save("children", &x.children)
}

func (x *frameRefnode) afterLoad() {}
func (x *frameRefnode) load(m state.Map) {
	m.Load("nrSegments", &x.nrSegments)
	m.Load("parent", &x.parent)
	m.Load("parentIndex", &x.parentIndex)
	m.Load("hasChildren", &x.hasChildren)
	m.Load("keys", &x.keys)
	m.Load("values", &x.values)
	m.Load("children", &x.children)
}

func (x *frameRefIterator) beforeSave() {}
func (x *frameRefIterator) save(m state.Map) {
	x.beforeSave()
	m.Save("node", &x.node)
	m.Save("index", &x.index)
}

func (x *frameRefIterator) afterLoad() {}
func (x *frameRefIterator) load(m state.Map) {
	m.Load("node", &x.node)
	m.Load("index", &x.index)
}

func (x *frameRefGapIterator) beforeSave() {}
func (x *frameRefGapIterator) save(m state.Map) {
	x.beforeSave()
	m.Save("node", &x.node)
	m.Save("index", &x.index)
}

func (x *frameRefGapIterator) afterLoad() {}
func (x *frameRefGapIterator) load(m state.Map) {
	m.Load("node", &x.node)
	m.Load("index", &x.index)
}

func (x *frameRefSegmentDataSlices) beforeSave() {}
func (x *frameRefSegmentDataSlices) save(m state.Map) {
	x.beforeSave()
	m.Save("Start", &x.Start)
	m.Save("End", &x.End)
	m.Save("Values", &x.Values)
}

func (x *frameRefSegmentDataSlices) afterLoad() {}
func (x *frameRefSegmentDataSlices) load(m state.Map) {
	m.Load("Start", &x.Start)
	m.Load("End", &x.End)
	m.Load("Values", &x.Values)
}

func (x *Handle) beforeSave() {}
func (x *Handle) save(m state.Map) {
	x.beforeSave()
	m.Save("HandleOperations", &x.HandleOperations)
	m.Save("dirCursor", &x.dirCursor)
}

func (x *Handle) afterLoad() {}
func (x *Handle) load(m state.Map) {
	m.Load("HandleOperations", &x.HandleOperations)
	m.Load("dirCursor", &x.dirCursor)
}

func (x *HostFileMapper) beforeSave() {}
func (x *HostFileMapper) save(m state.Map) {
	x.beforeSave()
	m.Save("refs", &x.refs)
}

func (x *HostFileMapper) load(m state.Map) {
	m.Load("refs", &x.refs)
	m.AfterLoad(x.afterLoad)
}

func (x *mapping) beforeSave() {}
func (x *mapping) save(m state.Map) {
	x.beforeSave()
	m.Save("addr", &x.addr)
	m.Save("writable", &x.writable)
}

func (x *mapping) afterLoad() {}
func (x *mapping) load(m state.Map) {
	m.Load("addr", &x.addr)
	m.Load("writable", &x.writable)
}

func (x *simpleInodeOperations) beforeSave() {}
func (x *simpleInodeOperations) save(m state.Map) {
	x.beforeSave()
	m.Save("InodeSimpleAttributes", &x.InodeSimpleAttributes)
}

func (x *simpleInodeOperations) afterLoad() {}
func (x *simpleInodeOperations) load(m state.Map) {
	m.Load("InodeSimpleAttributes", &x.InodeSimpleAttributes)
}

func (x *InodeSimpleAttributes) beforeSave() {}
func (x *InodeSimpleAttributes) save(m state.Map) {
	x.beforeSave()
	m.Save("FSType", &x.FSType)
	m.Save("UAttr", &x.UAttr)
}

func (x *InodeSimpleAttributes) afterLoad() {}
func (x *InodeSimpleAttributes) load(m state.Map) {
	m.Load("FSType", &x.FSType)
	m.Load("UAttr", &x.UAttr)
}

func (x *InMemoryAttributes) beforeSave() {}
func (x *InMemoryAttributes) save(m state.Map) {
	x.beforeSave()
	m.Save("Unstable", &x.Unstable)
	m.Save("Xattrs", &x.Xattrs)
}

func (x *InMemoryAttributes) afterLoad() {}
func (x *InMemoryAttributes) load(m state.Map) {
	m.Load("Unstable", &x.Unstable)
	m.Load("Xattrs", &x.Xattrs)
}

func (x *NoMappable) beforeSave() {}
func (x *NoMappable) save(m state.Map) {
	x.beforeSave()
}

func (x *NoMappable) afterLoad() {}
func (x *NoMappable) load(m state.Map) {
}

func (x *NoopWriteOut) beforeSave() {}
func (x *NoopWriteOut) save(m state.Map) {
	x.beforeSave()
}

func (x *NoopWriteOut) afterLoad() {}
func (x *NoopWriteOut) load(m state.Map) {
}

func (x *InodeNotDirectory) beforeSave() {}
func (x *InodeNotDirectory) save(m state.Map) {
	x.beforeSave()
}

func (x *InodeNotDirectory) afterLoad() {}
func (x *InodeNotDirectory) load(m state.Map) {
}

func (x *InodeNotSocket) beforeSave() {}
func (x *InodeNotSocket) save(m state.Map) {
	x.beforeSave()
}

func (x *InodeNotSocket) afterLoad() {}
func (x *InodeNotSocket) load(m state.Map) {
}

func (x *InodeNotRenameable) beforeSave() {}
func (x *InodeNotRenameable) save(m state.Map) {
	x.beforeSave()
}

func (x *InodeNotRenameable) afterLoad() {}
func (x *InodeNotRenameable) load(m state.Map) {
}

func (x *InodeNotOpenable) beforeSave() {}
func (x *InodeNotOpenable) save(m state.Map) {
	x.beforeSave()
}

func (x *InodeNotOpenable) afterLoad() {}
func (x *InodeNotOpenable) load(m state.Map) {
}

func (x *InodeNotVirtual) beforeSave() {}
func (x *InodeNotVirtual) save(m state.Map) {
	x.beforeSave()
}

func (x *InodeNotVirtual) afterLoad() {}
func (x *InodeNotVirtual) load(m state.Map) {
}

func (x *InodeNotSymlink) beforeSave() {}
func (x *InodeNotSymlink) save(m state.Map) {
	x.beforeSave()
}

func (x *InodeNotSymlink) afterLoad() {}
func (x *InodeNotSymlink) load(m state.Map) {
}

func (x *InodeNoExtendedAttributes) beforeSave() {}
func (x *InodeNoExtendedAttributes) save(m state.Map) {
	x.beforeSave()
}

func (x *InodeNoExtendedAttributes) afterLoad() {}
func (x *InodeNoExtendedAttributes) load(m state.Map) {
}

func (x *DeprecatedFileOperations) beforeSave() {}
func (x *DeprecatedFileOperations) save(m state.Map) {
	x.beforeSave()
}

func (x *DeprecatedFileOperations) afterLoad() {}
func (x *DeprecatedFileOperations) load(m state.Map) {
}

func (x *CachingInodeOperations) beforeSave() {}
func (x *CachingInodeOperations) save(m state.Map) {
	x.beforeSave()
	m.Save("backingFile", &x.backingFile)
	m.Save("platform", &x.platform)
	m.Save("forcePageCache", &x.forcePageCache)
	m.Save("attr", &x.attr)
	m.Save("dirtyAttr", &x.dirtyAttr)
	m.Save("mappings", &x.mappings)
	m.Save("cache", &x.cache)
	m.Save("dirty", &x.dirty)
	m.Save("hostFileMapper", &x.hostFileMapper)
	m.Save("refs", &x.refs)
}

func (x *CachingInodeOperations) afterLoad() {}
func (x *CachingInodeOperations) load(m state.Map) {
	m.Load("backingFile", &x.backingFile)
	m.Load("platform", &x.platform)
	m.Load("forcePageCache", &x.forcePageCache)
	m.Load("attr", &x.attr)
	m.Load("dirtyAttr", &x.dirtyAttr)
	m.Load("mappings", &x.mappings)
	m.Load("cache", &x.cache)
	m.Load("dirty", &x.dirty)
	m.Load("hostFileMapper", &x.hostFileMapper)
	m.Load("refs", &x.refs)
}

func (x *inodeReadWriter) beforeSave() {}
func (x *inodeReadWriter) save(m state.Map) {
	x.beforeSave()
	m.Save("ctx", &x.ctx)
	m.Save("c", &x.c)
	m.Save("offset", &x.offset)
}

func (x *inodeReadWriter) afterLoad() {}
func (x *inodeReadWriter) load(m state.Map) {
	m.Load("ctx", &x.ctx)
	m.Load("c", &x.c)
	m.Load("offset", &x.offset)
}

func init() {
	state.Register("fsutil.DirtySet", (*DirtySet)(nil), state.Fns{Save: (*DirtySet).save, Load: (*DirtySet).load})
	state.Register("fsutil.Dirtynode", (*Dirtynode)(nil), state.Fns{Save: (*Dirtynode).save, Load: (*Dirtynode).load})
	state.Register("fsutil.DirtyIterator", (*DirtyIterator)(nil), state.Fns{Save: (*DirtyIterator).save, Load: (*DirtyIterator).load})
	state.Register("fsutil.DirtyGapIterator", (*DirtyGapIterator)(nil), state.Fns{Save: (*DirtyGapIterator).save, Load: (*DirtyGapIterator).load})
	state.Register("fsutil.DirtySegmentDataSlices", (*DirtySegmentDataSlices)(nil), state.Fns{Save: (*DirtySegmentDataSlices).save, Load: (*DirtySegmentDataSlices).load})
	state.Register("fsutil.NoopRelease", (*NoopRelease)(nil), state.Fns{Save: (*NoopRelease).save, Load: (*NoopRelease).load})
	state.Register("fsutil.GenericSeek", (*GenericSeek)(nil), state.Fns{Save: (*GenericSeek).save, Load: (*GenericSeek).load})
	state.Register("fsutil.ZeroSeek", (*ZeroSeek)(nil), state.Fns{Save: (*ZeroSeek).save, Load: (*ZeroSeek).load})
	state.Register("fsutil.PipeSeek", (*PipeSeek)(nil), state.Fns{Save: (*PipeSeek).save, Load: (*PipeSeek).load})
	state.Register("fsutil.NotDirReaddir", (*NotDirReaddir)(nil), state.Fns{Save: (*NotDirReaddir).save, Load: (*NotDirReaddir).load})
	state.Register("fsutil.NoFsync", (*NoFsync)(nil), state.Fns{Save: (*NoFsync).save, Load: (*NoFsync).load})
	state.Register("fsutil.NoopFsync", (*NoopFsync)(nil), state.Fns{Save: (*NoopFsync).save, Load: (*NoopFsync).load})
	state.Register("fsutil.NoopFlush", (*NoopFlush)(nil), state.Fns{Save: (*NoopFlush).save, Load: (*NoopFlush).load})
	state.Register("fsutil.NoMMap", (*NoMMap)(nil), state.Fns{Save: (*NoMMap).save, Load: (*NoMMap).load})
	state.Register("fsutil.NoIoctl", (*NoIoctl)(nil), state.Fns{Save: (*NoIoctl).save, Load: (*NoIoctl).load})
	state.Register("fsutil.DirFileOperations", (*DirFileOperations)(nil), state.Fns{Save: (*DirFileOperations).save, Load: (*DirFileOperations).load})
	state.Register("fsutil.FileRangeSet", (*FileRangeSet)(nil), state.Fns{Save: (*FileRangeSet).save, Load: (*FileRangeSet).load})
	state.Register("fsutil.FileRangenode", (*FileRangenode)(nil), state.Fns{Save: (*FileRangenode).save, Load: (*FileRangenode).load})
	state.Register("fsutil.FileRangeIterator", (*FileRangeIterator)(nil), state.Fns{Save: (*FileRangeIterator).save, Load: (*FileRangeIterator).load})
	state.Register("fsutil.FileRangeGapIterator", (*FileRangeGapIterator)(nil), state.Fns{Save: (*FileRangeGapIterator).save, Load: (*FileRangeGapIterator).load})
	state.Register("fsutil.FileRangeSegmentDataSlices", (*FileRangeSegmentDataSlices)(nil), state.Fns{Save: (*FileRangeSegmentDataSlices).save, Load: (*FileRangeSegmentDataSlices).load})
	state.Register("fsutil.frameRefSet", (*frameRefSet)(nil), state.Fns{Save: (*frameRefSet).save, Load: (*frameRefSet).load})
	state.Register("fsutil.frameRefnode", (*frameRefnode)(nil), state.Fns{Save: (*frameRefnode).save, Load: (*frameRefnode).load})
	state.Register("fsutil.frameRefIterator", (*frameRefIterator)(nil), state.Fns{Save: (*frameRefIterator).save, Load: (*frameRefIterator).load})
	state.Register("fsutil.frameRefGapIterator", (*frameRefGapIterator)(nil), state.Fns{Save: (*frameRefGapIterator).save, Load: (*frameRefGapIterator).load})
	state.Register("fsutil.frameRefSegmentDataSlices", (*frameRefSegmentDataSlices)(nil), state.Fns{Save: (*frameRefSegmentDataSlices).save, Load: (*frameRefSegmentDataSlices).load})
	state.Register("fsutil.Handle", (*Handle)(nil), state.Fns{Save: (*Handle).save, Load: (*Handle).load})
	state.Register("fsutil.HostFileMapper", (*HostFileMapper)(nil), state.Fns{Save: (*HostFileMapper).save, Load: (*HostFileMapper).load})
	state.Register("fsutil.mapping", (*mapping)(nil), state.Fns{Save: (*mapping).save, Load: (*mapping).load})
	state.Register("fsutil.simpleInodeOperations", (*simpleInodeOperations)(nil), state.Fns{Save: (*simpleInodeOperations).save, Load: (*simpleInodeOperations).load})
	state.Register("fsutil.InodeSimpleAttributes", (*InodeSimpleAttributes)(nil), state.Fns{Save: (*InodeSimpleAttributes).save, Load: (*InodeSimpleAttributes).load})
	state.Register("fsutil.InMemoryAttributes", (*InMemoryAttributes)(nil), state.Fns{Save: (*InMemoryAttributes).save, Load: (*InMemoryAttributes).load})
	state.Register("fsutil.NoMappable", (*NoMappable)(nil), state.Fns{Save: (*NoMappable).save, Load: (*NoMappable).load})
	state.Register("fsutil.NoopWriteOut", (*NoopWriteOut)(nil), state.Fns{Save: (*NoopWriteOut).save, Load: (*NoopWriteOut).load})
	state.Register("fsutil.InodeNotDirectory", (*InodeNotDirectory)(nil), state.Fns{Save: (*InodeNotDirectory).save, Load: (*InodeNotDirectory).load})
	state.Register("fsutil.InodeNotSocket", (*InodeNotSocket)(nil), state.Fns{Save: (*InodeNotSocket).save, Load: (*InodeNotSocket).load})
	state.Register("fsutil.InodeNotRenameable", (*InodeNotRenameable)(nil), state.Fns{Save: (*InodeNotRenameable).save, Load: (*InodeNotRenameable).load})
	state.Register("fsutil.InodeNotOpenable", (*InodeNotOpenable)(nil), state.Fns{Save: (*InodeNotOpenable).save, Load: (*InodeNotOpenable).load})
	state.Register("fsutil.InodeNotVirtual", (*InodeNotVirtual)(nil), state.Fns{Save: (*InodeNotVirtual).save, Load: (*InodeNotVirtual).load})
	state.Register("fsutil.InodeNotSymlink", (*InodeNotSymlink)(nil), state.Fns{Save: (*InodeNotSymlink).save, Load: (*InodeNotSymlink).load})
	state.Register("fsutil.InodeNoExtendedAttributes", (*InodeNoExtendedAttributes)(nil), state.Fns{Save: (*InodeNoExtendedAttributes).save, Load: (*InodeNoExtendedAttributes).load})
	state.Register("fsutil.DeprecatedFileOperations", (*DeprecatedFileOperations)(nil), state.Fns{Save: (*DeprecatedFileOperations).save, Load: (*DeprecatedFileOperations).load})
	state.Register("fsutil.CachingInodeOperations", (*CachingInodeOperations)(nil), state.Fns{Save: (*CachingInodeOperations).save, Load: (*CachingInodeOperations).load})
	state.Register("fsutil.inodeReadWriter", (*inodeReadWriter)(nil), state.Fns{Save: (*inodeReadWriter).save, Load: (*inodeReadWriter).load})
}