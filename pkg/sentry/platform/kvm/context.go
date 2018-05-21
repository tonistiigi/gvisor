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

package kvm

import (
	"gvisor.googlesource.com/gvisor/pkg/sentry/arch"
	"gvisor.googlesource.com/gvisor/pkg/sentry/platform"
	"gvisor.googlesource.com/gvisor/pkg/sentry/platform/interrupt"
	"gvisor.googlesource.com/gvisor/pkg/sentry/platform/ring0"
	"gvisor.googlesource.com/gvisor/pkg/sentry/usermem"
)

// context is an implementation of the platform context.
//
// This is a thin wrapper around the machine.
type context struct {
	// machine is the parent machine, and is immutable.
	machine *machine

	// interrupt is the interrupt context.
	interrupt interrupt.Forwarder
}

// Switch runs the provided context in the given address space.
func (c *context) Switch(as platform.AddressSpace, ac arch.Context, _ int32) (*arch.SignalInfo, usermem.AccessType, error) {
	// Extract data.
	localAS := as.(*addressSpace)
	regs := &ac.StateData().Regs
	fp := (*byte)(ac.FloatingPointData())

	// Grab a vCPU.
	cpu := c.machine.Get()

	// Enable interrupts (i.e. calls to vCPU.Notify).
	if !c.interrupt.Enable(cpu) {
		c.machine.Put(cpu) // Already preempted.
		return nil, usermem.NoAccess, platform.ErrContextInterrupt
	}

	// Set the active address space.
	//
	// This must be done prior to the call to Touch below. If the address
	// space is invalidated between this line and the call below, we will
	// flag on entry anyways. When the active address space below is
	// cleared, it indicates that we don't need an explicit interrupt and
	// that the flush can occur naturally on the next user entry.
	cpu.active.set(localAS)

	// Mark the address space as dirty.
	flags := ring0.Flags(0)
	if localAS.Touch(cpu) {
		flags |= ring0.FlagFlush
	}
	if ac.FullRestore() {
		flags |= ring0.FlagFull
	}

	// Take the blue pill.
	si, at, err := cpu.SwitchToUser(regs, fp, localAS.pageTables, flags)

	// Clear the address space.
	cpu.active.set(nil)

	// Release resources.
	c.machine.Put(cpu)

	// All done.
	c.interrupt.Disable()
	return si, at, err
}

// Interrupt interrupts the running context.
func (c *context) Interrupt() {
	c.interrupt.NotifyInterrupt()
}
