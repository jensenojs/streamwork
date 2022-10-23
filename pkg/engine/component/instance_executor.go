package component

import (
	"streamwork/pkg/engine"
)

// =================================================================
// implement for InstanceExecutor
func (i *InstanceExecutorImpl) SetIncoming(in *engine.EventQueue) {
	i.Incoming = in
}

func (i *InstanceExecutorImpl) SetOutgoing(out *engine.EventQueue) {
	i.Outgoing = out
}

// helper functions to receive events
func (i *InstanceExecutorImpl) TakeIncomingEvent() engine.Event {
	e, ok := <-i.Incoming.Queue
	if ok {
		return e
	}
	return nil
}

// helper functions to send events
func (i *InstanceExecutorImpl) SendOutgoingEvent(event engine.Event) {
	i.Outgoing.Queue <- event
}

// =================================================================
// implement for Process
func (i *InstanceExecutorImpl) NewProcess() {
	i.FnWrapper = func() {
		for {
			if ok := i.Fn(); ok != true {
				break
			}
		}
	}
}

func (i *InstanceExecutorImpl) Start() {
	go i.FnWrapper()
}

func (i *InstanceExecutorImpl) RunOnce() bool {
	panic("Need specific implementation")
}

// helper function to set RunOnce
func (i *InstanceExecutorImpl) SetRunOnce(RunOnce func() bool) {
	i.Fn = RunOnce
}
