// automatically generated by stateify.

package ping

import (
	"gvisor.googlesource.com/gvisor/pkg/state"
	"gvisor.googlesource.com/gvisor/pkg/tcpip/buffer"
)

func (x *pingPacket) beforeSave() {}
func (x *pingPacket) save(m state.Map) {
	x.beforeSave()
	var data buffer.VectorisedView = x.saveData()
	m.SaveValue("data", data)
	m.Save("pingPacketEntry", &x.pingPacketEntry)
	m.Save("senderAddress", &x.senderAddress)
	m.Save("timestamp", &x.timestamp)
	m.Save("hasTimestamp", &x.hasTimestamp)
}

func (x *pingPacket) afterLoad() {}
func (x *pingPacket) load(m state.Map) {
	m.Load("pingPacketEntry", &x.pingPacketEntry)
	m.Load("senderAddress", &x.senderAddress)
	m.Load("timestamp", &x.timestamp)
	m.Load("hasTimestamp", &x.hasTimestamp)
	m.LoadValue("data", new(buffer.VectorisedView), func(y interface{}) { x.loadData(y.(buffer.VectorisedView)) })
}

func (x *pingPacketList) beforeSave() {}
func (x *pingPacketList) save(m state.Map) {
	x.beforeSave()
	m.Save("head", &x.head)
	m.Save("tail", &x.tail)
}

func (x *pingPacketList) afterLoad() {}
func (x *pingPacketList) load(m state.Map) {
	m.Load("head", &x.head)
	m.Load("tail", &x.tail)
}

func (x *pingPacketEntry) beforeSave() {}
func (x *pingPacketEntry) save(m state.Map) {
	x.beforeSave()
	m.Save("next", &x.next)
	m.Save("prev", &x.prev)
}

func (x *pingPacketEntry) afterLoad() {}
func (x *pingPacketEntry) load(m state.Map) {
	m.Load("next", &x.next)
	m.Load("prev", &x.prev)
}

func init() {
	state.Register("ping.pingPacket", (*pingPacket)(nil), state.Fns{Save: (*pingPacket).save, Load: (*pingPacket).load})
	state.Register("ping.pingPacketList", (*pingPacketList)(nil), state.Fns{Save: (*pingPacketList).save, Load: (*pingPacketList).load})
	state.Register("ping.pingPacketEntry", (*pingPacketEntry)(nil), state.Fns{Save: (*pingPacketEntry).save, Load: (*pingPacketEntry).load})
}