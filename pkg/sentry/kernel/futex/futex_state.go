// automatically generated by stateify.

package futex

import (
	"gvisor.googlesource.com/gvisor/pkg/state"
)

func (x *Waiter) beforeSave() {}
func (x *Waiter) save(m state.Map) {
	x.beforeSave()
	m.Save("waiterEntry", &x.waiterEntry)
	m.Save("complete", &x.complete)
	m.Save("C", &x.C)
	m.Save("addr", &x.addr)
	m.Save("bitmask", &x.bitmask)
}

func (x *Waiter) afterLoad() {}
func (x *Waiter) load(m state.Map) {
	m.Load("waiterEntry", &x.waiterEntry)
	m.Load("complete", &x.complete)
	m.Load("C", &x.C)
	m.Load("addr", &x.addr)
	m.Load("bitmask", &x.bitmask)
}

func (x *bucket) beforeSave() {}
func (x *bucket) save(m state.Map) {
	x.beforeSave()
	if !state.IsZeroValue(x.waiters) { m.Failf("waiters is %v, expected zero", x.waiters) }
}

func (x *bucket) afterLoad() {}
func (x *bucket) load(m state.Map) {
}

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
	state.Register("futex.Waiter", (*Waiter)(nil), state.Fns{Save: (*Waiter).save, Load: (*Waiter).load})
	state.Register("futex.bucket", (*bucket)(nil), state.Fns{Save: (*bucket).save, Load: (*bucket).load})
	state.Register("futex.Manager", (*Manager)(nil), state.Fns{Save: (*Manager).save, Load: (*Manager).load})
	state.Register("futex.waiterList", (*waiterList)(nil), state.Fns{Save: (*waiterList).save, Load: (*waiterList).load})
	state.Register("futex.waiterEntry", (*waiterEntry)(nil), state.Fns{Save: (*waiterEntry).save, Load: (*waiterEntry).load})
}
