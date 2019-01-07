// automatically generated by stateify.

package kvm

import (
	"gvisor.googlesource.com/gvisor/pkg/state"
)

func (x *hostMapSet) beforeSave() {}
func (x *hostMapSet) save(m state.Map) {
	x.beforeSave()
	var root *hostMapSegmentDataSlices = x.saveRoot()
	m.SaveValue("root", root)
}

func (x *hostMapSet) afterLoad() {}
func (x *hostMapSet) load(m state.Map) {
	m.LoadValue("root", new(*hostMapSegmentDataSlices), func(y interface{}) { x.loadRoot(y.(*hostMapSegmentDataSlices)) })
}

func (x *hostMapnode) beforeSave() {}
func (x *hostMapnode) save(m state.Map) {
	x.beforeSave()
	m.Save("nrSegments", &x.nrSegments)
	m.Save("parent", &x.parent)
	m.Save("parentIndex", &x.parentIndex)
	m.Save("hasChildren", &x.hasChildren)
	m.Save("keys", &x.keys)
	m.Save("values", &x.values)
	m.Save("children", &x.children)
}

func (x *hostMapnode) afterLoad() {}
func (x *hostMapnode) load(m state.Map) {
	m.Load("nrSegments", &x.nrSegments)
	m.Load("parent", &x.parent)
	m.Load("parentIndex", &x.parentIndex)
	m.Load("hasChildren", &x.hasChildren)
	m.Load("keys", &x.keys)
	m.Load("values", &x.values)
	m.Load("children", &x.children)
}

func (x *hostMapSegmentDataSlices) beforeSave() {}
func (x *hostMapSegmentDataSlices) save(m state.Map) {
	x.beforeSave()
	m.Save("Start", &x.Start)
	m.Save("End", &x.End)
	m.Save("Values", &x.Values)
}

func (x *hostMapSegmentDataSlices) afterLoad() {}
func (x *hostMapSegmentDataSlices) load(m state.Map) {
	m.Load("Start", &x.Start)
	m.Load("End", &x.End)
	m.Load("Values", &x.Values)
}

func init() {
	state.Register("kvm.hostMapSet", (*hostMapSet)(nil), state.Fns{Save: (*hostMapSet).save, Load: (*hostMapSet).load})
	state.Register("kvm.hostMapnode", (*hostMapnode)(nil), state.Fns{Save: (*hostMapnode).save, Load: (*hostMapnode).load})
	state.Register("kvm.hostMapSegmentDataSlices", (*hostMapSegmentDataSlices)(nil), state.Fns{Save: (*hostMapSegmentDataSlices).save, Load: (*hostMapSegmentDataSlices).load})
}