// automatically generated by stateify.

package futex

import (
	"gvisor.googlesource.com/gvisor/pkg/state"
)

func (x *Manager) beforeSave() {}
func (x *Manager) save(m state.Map) {
	x.beforeSave()
	if !state.IsZeroValue(x.buckets) { m.Failf("buckets is %v, expected zero", x.buckets) }
}

func (x *Manager) afterLoad() {}
func (x *Manager) load(m state.Map) {
}

func (x *waiterList) beforeSave() {}
func (x *waiterList) save(m state.Map) {
	x.beforeSave()
	m.Save("head", &x.head)
	m.Save("tail", &x.tail)
}

func (x *waiterList) afterLoad() {}
func (x *waiterList) load(m state.Map) {
	m.Load("head", &x.head)
	m.Load("tail", &x.tail)
}

func (x *waiterEntry) beforeSave() {}
func (x *waiterEntry) save(m state.Map) {
	x.beforeSave()
	m.Save("next", &x.next)
	m.Save("prev", &x.prev)
}

func (x *waiterEntry) afterLoad() {}
func (x *waiterEntry) load(m state.Map) {
	m.Load("next", &x.next)
	m.Load("prev", &x.prev)
}

func init() {
	state.Register("futex.Manager", (*Manager)(nil), state.Fns{Save: (*Manager).save, Load: (*Manager).load})
	state.Register("futex.waiterList", (*waiterList)(nil), state.Fns{Save: (*waiterList).save, Load: (*waiterList).load})
	state.Register("futex.waiterEntry", (*waiterEntry)(nil), state.Fns{Save: (*waiterEntry).save, Load: (*waiterEntry).load})
}