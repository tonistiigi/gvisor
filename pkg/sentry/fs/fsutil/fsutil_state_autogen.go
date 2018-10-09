// automatically generated by stateify.

package fsutil

import (
	"gvisor.googlesource.com/gvisor/pkg/state"
)

func (x *DirtyInfo) beforeSave() {}
func (x *DirtyInfo) save(m state.Map) {
	x.beforeSave()
	m.Save("Keep", &x.Keep)
}

func (x *DirtyInfo) afterLoad() {}
func (x *DirtyInfo) load(m state.Map) {
	m.Load("Keep", &x.Keep)
}

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

func init() {
	state.Register("fsutil.DirtyInfo", (*DirtyInfo)(nil), state.Fns{Save: (*DirtyInfo).save, Load: (*DirtyInfo).load})
	state.Register("fsutil.DirtySet", (*DirtySet)(nil), state.Fns{Save: (*DirtySet).save, Load: (*DirtySet).load})
	state.Register("fsutil.Dirtynode", (*Dirtynode)(nil), state.Fns{Save: (*Dirtynode).save, Load: (*Dirtynode).load})
	state.Register("fsutil.DirtySegmentDataSlices", (*DirtySegmentDataSlices)(nil), state.Fns{Save: (*DirtySegmentDataSlices).save, Load: (*DirtySegmentDataSlices).load})
	state.Register("fsutil.DirFileOperations", (*DirFileOperations)(nil), state.Fns{Save: (*DirFileOperations).save, Load: (*DirFileOperations).load})
	state.Register("fsutil.FileRangeSet", (*FileRangeSet)(nil), state.Fns{Save: (*FileRangeSet).save, Load: (*FileRangeSet).load})
	state.Register("fsutil.FileRangenode", (*FileRangenode)(nil), state.Fns{Save: (*FileRangenode).save, Load: (*FileRangenode).load})
	state.Register("fsutil.FileRangeSegmentDataSlices", (*FileRangeSegmentDataSlices)(nil), state.Fns{Save: (*FileRangeSegmentDataSlices).save, Load: (*FileRangeSegmentDataSlices).load})
	state.Register("fsutil.frameRefSet", (*frameRefSet)(nil), state.Fns{Save: (*frameRefSet).save, Load: (*frameRefSet).load})
	state.Register("fsutil.frameRefnode", (*frameRefnode)(nil), state.Fns{Save: (*frameRefnode).save, Load: (*frameRefnode).load})
	state.Register("fsutil.frameRefSegmentDataSlices", (*frameRefSegmentDataSlices)(nil), state.Fns{Save: (*frameRefSegmentDataSlices).save, Load: (*frameRefSegmentDataSlices).load})
	state.Register("fsutil.Handle", (*Handle)(nil), state.Fns{Save: (*Handle).save, Load: (*Handle).load})
	state.Register("fsutil.HostFileMapper", (*HostFileMapper)(nil), state.Fns{Save: (*HostFileMapper).save, Load: (*HostFileMapper).load})
	state.Register("fsutil.simpleInodeOperations", (*simpleInodeOperations)(nil), state.Fns{Save: (*simpleInodeOperations).save, Load: (*simpleInodeOperations).load})
	state.Register("fsutil.InodeSimpleAttributes", (*InodeSimpleAttributes)(nil), state.Fns{Save: (*InodeSimpleAttributes).save, Load: (*InodeSimpleAttributes).load})
	state.Register("fsutil.InMemoryAttributes", (*InMemoryAttributes)(nil), state.Fns{Save: (*InMemoryAttributes).save, Load: (*InMemoryAttributes).load})
	state.Register("fsutil.CachingInodeOperations", (*CachingInodeOperations)(nil), state.Fns{Save: (*CachingInodeOperations).save, Load: (*CachingInodeOperations).load})
}