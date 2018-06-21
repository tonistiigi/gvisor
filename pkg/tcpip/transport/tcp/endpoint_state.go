// Copyright 2017 The Netstack Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tcp

import (
	"fmt"
	"sync"

	"gvisor.googlesource.com/gvisor/pkg/tcpip"
	"gvisor.googlesource.com/gvisor/pkg/tcpip/header"
	"gvisor.googlesource.com/gvisor/pkg/tcpip/stack"
)

func (e *endpoint) drainSegmentLocked() {
	// Drain only up to once.
	if e.drainDone != nil {
		return
	}

	e.drainDone = make(chan struct{})
	e.undrain = make(chan struct{})
	e.mu.Unlock()

	e.notifyProtocolGoroutine(notifyDrain)
	<-e.drainDone

	e.mu.Lock()
}

// beforeSave is invoked by stateify.
func (e *endpoint) beforeSave() {
	// Stop incoming packets.
	e.segmentQueue.setLimit(0)

	e.mu.Lock()
	defer e.mu.Unlock()

	switch e.state {
	case stateInitial, stateBound:
	case stateListen, stateConnecting, stateConnected:
		if e.state == stateConnected && !e.workerRunning {
			// The endpoint must be in acceptedChan.
			break
		}
		e.drainSegmentLocked()
		if e.state != stateClosed && e.state != stateError {
			if !e.workerRunning {
				panic("endpoint has no worker running in listen, connecting, or connected state")
			}
			break
		}
		fallthrough
	case stateClosed, stateError:
		if e.workerRunning {
			panic("endpoint still has worker running in closed or error state")
		}
	default:
		panic(fmt.Sprintf("endpoint in unknown state %v", e.state))
	}

	if e.waiterQueue != nil && !e.waiterQueue.IsEmpty() {
		panic("endpoint still has waiters upon save")
	}

	if !((e.state == stateBound || e.state == stateListen) == e.isPortReserved) {
		panic("endpoint port must and must only be reserved in bound or listen state")
	}

	if e.acceptedChan != nil {
		close(e.acceptedChan)
		e.acceptedEndpoints = make([]*endpoint, len(e.acceptedChan), cap(e.acceptedChan))
		i := 0
		for ep := range e.acceptedChan {
			e.acceptedEndpoints[i] = ep
			i++
		}
		if i != len(e.acceptedEndpoints) {
			panic("endpoint acceptedChan buffer got consumed by background context")
		}
	}
}

// saveState is invoked by stateify.
func (e *endpoint) saveState() endpointState {
	return e.state
}

// Endpoint loading must be done in the following ordering by their state, to
// avoid dangling connecting w/o listening peer, and to avoid conflicts in port
// reservation.
var connectedLoading sync.WaitGroup
var listenLoading sync.WaitGroup
var connectingLoading sync.WaitGroup

// Bound endpoint loading happens last.

// loadState is invoked by stateify.
func (e *endpoint) loadState(state endpointState) {
	// This is to ensure that the loading wait groups include all applicable
	// endpoints before any asynchronous calls to the Wait() methods.
	switch state {
	case stateConnected:
		connectedLoading.Add(1)
	case stateListen:
		listenLoading.Add(1)
	case stateConnecting:
		connectingLoading.Add(1)
	}
	e.state = state
}

// afterLoad is invoked by stateify.
func (e *endpoint) afterLoad() {
	// We load acceptedChan buffer indirectly here. Note that closed
	// endpoints might not need to allocate the channel.
	// FIXME
	if cap(e.acceptedEndpoints) > 0 {
		e.acceptedChan = make(chan *endpoint, cap(e.acceptedEndpoints))
		for _, ep := range e.acceptedEndpoints {
			e.acceptedChan <- ep
		}
		e.acceptedEndpoints = nil
	}

	e.stack = stack.StackFromEnv
	e.segmentQueue.setLimit(2 * e.rcvBufSize)
	e.workMu.Init()

	state := e.state
	switch state {
	case stateInitial, stateBound, stateListen, stateConnecting, stateConnected:
		var ss SendBufferSizeOption
		if err := e.stack.TransportProtocolOption(ProtocolNumber, &ss); err == nil {
			if e.sndBufSize < ss.Min || e.sndBufSize > ss.Max {
				panic(fmt.Sprintf("endpoint.sndBufSize %d is outside the min and max allowed [%d, %d]", e.sndBufSize, ss.Min, ss.Max))
			}
			if e.rcvBufSize < ss.Min || e.rcvBufSize > ss.Max {
				panic(fmt.Sprintf("endpoint.rcvBufSize %d is outside the min and max allowed [%d, %d]", e.rcvBufSize, ss.Min, ss.Max))
			}
		}
	}

	bind := func() {
		e.state = stateInitial
		if len(e.bindAddress) == 0 {
			e.bindAddress = e.id.LocalAddress
		}
		if err := e.Bind(tcpip.FullAddress{Addr: e.bindAddress, Port: e.id.LocalPort}, nil); err != nil {
			panic("endpoint binding failed: " + err.String())
		}
	}

	switch state {
	case stateConnected:
		bind()
		if len(e.connectingAddress) == 0 {
			// This endpoint is accepted by netstack but not yet by
			// the app. If the endpoint is IPv6 but the remote
			// address is IPv4, we need to connect as IPv6 so that
			// dual-stack mode can be properly activated.
			if e.netProto == header.IPv6ProtocolNumber && len(e.id.RemoteAddress) != header.IPv6AddressSize {
				e.connectingAddress = "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xff\xff" + e.id.RemoteAddress
			} else {
				e.connectingAddress = e.id.RemoteAddress
			}
		}
		if err := e.connect(tcpip.FullAddress{NIC: e.boundNICID, Addr: e.connectingAddress, Port: e.id.RemotePort}, false, e.workerRunning); err != tcpip.ErrConnectStarted {
			panic("endpoint connecting failed: " + err.String())
		}
		connectedLoading.Done()
	case stateListen:
		tcpip.AsyncLoading.Add(1)
		go func() {
			connectedLoading.Wait()
			bind()
			backlog := cap(e.acceptedChan)
			if err := e.Listen(backlog); err != nil {
				panic("endpoint listening failed: " + err.String())
			}
			listenLoading.Done()
			tcpip.AsyncLoading.Done()
		}()
	case stateConnecting:
		tcpip.AsyncLoading.Add(1)
		go func() {
			connectedLoading.Wait()
			listenLoading.Wait()
			bind()
			if err := e.Connect(tcpip.FullAddress{NIC: e.boundNICID, Addr: e.connectingAddress, Port: e.id.RemotePort}); err != tcpip.ErrConnectStarted {
				panic("endpoint connecting failed: " + err.String())
			}
			connectingLoading.Done()
			tcpip.AsyncLoading.Done()
		}()
	case stateBound:
		tcpip.AsyncLoading.Add(1)
		go func() {
			connectedLoading.Wait()
			listenLoading.Wait()
			connectingLoading.Wait()
			bind()
			tcpip.AsyncLoading.Done()
		}()
	case stateClosed, stateError:
		tcpip.DeleteDanglingEndpoint(e)
	}
}

// saveLastError is invoked by stateify.
func (e *endpoint) saveLastError() string {
	if e.lastError == nil {
		return ""
	}

	return e.lastError.String()
}

// loadLastError is invoked by stateify.
func (e *endpoint) loadLastError(s string) {
	if s == "" {
		return
	}

	e.lastError = loadError(s)
}

// saveHardError is invoked by stateify.
func (e *endpoint) saveHardError() string {
	if e.hardError == nil {
		return ""
	}

	return e.hardError.String()
}

// loadHardError is invoked by stateify.
func (e *endpoint) loadHardError(s string) {
	if s == "" {
		return
	}

	e.hardError = loadError(s)
}

var messageToError map[string]*tcpip.Error

var populate sync.Once

func loadError(s string) *tcpip.Error {
	populate.Do(func() {
		var errors = []*tcpip.Error{
			tcpip.ErrUnknownProtocol,
			tcpip.ErrUnknownNICID,
			tcpip.ErrUnknownProtocolOption,
			tcpip.ErrDuplicateNICID,
			tcpip.ErrDuplicateAddress,
			tcpip.ErrNoRoute,
			tcpip.ErrBadLinkEndpoint,
			tcpip.ErrAlreadyBound,
			tcpip.ErrInvalidEndpointState,
			tcpip.ErrAlreadyConnecting,
			tcpip.ErrAlreadyConnected,
			tcpip.ErrNoPortAvailable,
			tcpip.ErrPortInUse,
			tcpip.ErrBadLocalAddress,
			tcpip.ErrClosedForSend,
			tcpip.ErrClosedForReceive,
			tcpip.ErrWouldBlock,
			tcpip.ErrConnectionRefused,
			tcpip.ErrTimeout,
			tcpip.ErrAborted,
			tcpip.ErrConnectStarted,
			tcpip.ErrDestinationRequired,
			tcpip.ErrNotSupported,
			tcpip.ErrQueueSizeNotSupported,
			tcpip.ErrNotConnected,
			tcpip.ErrConnectionReset,
			tcpip.ErrConnectionAborted,
			tcpip.ErrNoSuchFile,
			tcpip.ErrInvalidOptionValue,
			tcpip.ErrNoLinkAddress,
			tcpip.ErrBadAddress,
			tcpip.ErrNetworkUnreachable,
		}

		messageToError = make(map[string]*tcpip.Error)
		for _, e := range errors {
			if messageToError[e.String()] != nil {
				panic("tcpip errors with duplicated message: " + e.String())
			}
			messageToError[e.String()] = e
		}
	})

	e, ok := messageToError[s]
	if !ok {
		panic("unknown error message: " + s)
	}

	return e
}
