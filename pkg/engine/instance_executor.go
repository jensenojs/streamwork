package engine

import "streamwork/pkg/api"

/**
 * InstanceExecutor take charge of componentExecutor in v0.1, as implementation of parallelism mechanism
 */
type InstanceExecutor interface {
	Process
	api.Component

	SetIncomingQueue(in *EventQueue)

	SetOutgoingQueue(out *EventQueue)
}

type InstanceExecutorImpl struct {
	name           string
	fnWrapper      func() 	   // wrapper function for fn
	fn             func() bool // process function, need to specific implementation for user logic
	stream         *api.Stream // connect to next component
	eventCollector []api.Event // accept events from user logic
	incomingQueue  *EventQueue // for upstream processes
	outgoingQueue  *EventQueue // for downstream processes
}

// =================================================================
// implement for Component
func (i *InstanceExecutorImpl) SetName(name string) { // Use InitNameAndStream to instead
	i.name = name
}

func (i *InstanceExecutorImpl) GetName() string {
	return i.name
}

func (i *InstanceExecutorImpl) SetOutgoingStream() { // Use InitNameAndStream to instead
	if i.stream == nil {
		i.stream = api.NewStream()
	}
}

func (i *InstanceExecutorImpl) GetOutgoingStream() *api.Stream {
	return i.stream
}

// helper function to init a instance executor combine SetName or SetOutgoingStream
func (i *InstanceExecutorImpl) InitNameAndStream(name string) {
	i.SetName(name)
	i.SetOutgoingStream()
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
