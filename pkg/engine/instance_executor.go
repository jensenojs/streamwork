package engine

import (
	"streamwork/pkg/api"
)

/**
 * Due to the need to achieve concurrency,
 * InstanceExecutor takes on some of ComponentExecutor's responsibilities in v0.1
 */
type InstanceExecutor interface {
	Process

	SetIncomingQueue(in *EventQueue)

	SetOutgoingQueue(out *EventQueue)
}

type InstanceExecutorImpl struct {
	fnWrapper      func()      // wrapper function for fn
	fn             func() bool // process function, need to specific implementation for user logic
	eventCollector []api.Event // accept events from user logic
	incomingQueue  *EventQueue // for upstream processes
	outgoingQueue  *EventQueue // for downstream processes
}

// =================================================================
// implement for InstanceExecutor
func (i *InstanceExecutorImpl) SetIncomingQueue(in *EventQueue) {
	i.incomingQueue = in
}

func (i *InstanceExecutorImpl) SetOutgoingQueue(out *EventQueue) {
	i.outgoingQueue = out
}

// helper functions to receive events
func (i *InstanceExecutorImpl) takeIncomingEvent() api.Event {
	e, ok := <-i.incomingQueue.queue
	if ok {
		return e
	}
	return nil
}

// helper functions to send events
func (i *InstanceExecutorImpl) sendOutgoingEvent(event api.Event) {
	i.outgoingQueue.queue <- event
}

// =================================================================
// implement for Process
func (i *InstanceExecutorImpl) newProcess() {
	i.fnWrapper = func() {
		for {
			if ok := i.fn(); ok != true {
				break
			}
		}
	}
}

func (i *InstanceExecutorImpl) Start() {
	go i.fnWrapper()
}

func (i *InstanceExecutorImpl) runOnce() bool {
	panic("Need specific implementation")
}

// helper function to set runOnce
func (i *InstanceExecutorImpl) setRunOnce(runOnce func() bool) {
	i.fn = runOnce
}
