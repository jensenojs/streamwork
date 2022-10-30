package transport

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/operator"
)

func NewEventDispatcher(downStreamExecutor *operator.OperatorExecutor) *EventDispatcher {
	return &EventDispatcher{
		downStreamExecutor: downStreamExecutor,
	}
}

// ============================================================
// implementation of Process

func (v *EventDispatcher) NewProcess() {
	v.fnWrapper = func() {
		for {
			if ok := v.RunOnce(); ok != true {
				break
			}
		}
	}
}

func (v *EventDispatcher) Start() {
	go v.fnWrapper()
}

func (v *EventDispatcher) RunOnce() bool {
	e := v.takeIncomingEvent()
	s := v.downStreamExecutor.GetGroupingStrategy()
	idx := s.GetInstance(e, len(v.outgoings))
	v.sendOutgoingEvent(e, idx)
	return true
}

// ============================================================

// init and use for incoming/outgoing queues
func (v *EventDispatcher) takeIncomingEvent() engine.Event {
	return v.incoming.Take()
}

func (v *EventDispatcher) sendOutgoingEvent(event engine.Event, idx int) {
	v.outgoings[idx].Send(event)
}

func (v *EventDispatcher) SetIncoming(queue engine.EventQueue) {
	v.incoming = queue
}

func (v *EventDispatcher) SetOutgoings(queues []engine.EventQueue) {
	v.outgoings = queues
}
