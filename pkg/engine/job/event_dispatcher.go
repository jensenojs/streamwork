package job

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/operator"
)

// EventDispatcher is responsible for transporting events from
// the incoming queue to the outgoing queues with a grouping strategy.
type EventDispatcher struct {
	fnWrapper          func() // wrapper function for fn, no need for fn
	downStreamExecutor *operator.OperatorExecutor
	incoming           *engine.EventQueue
	outgoings          []*engine.EventQueue
}

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
	e, ok := <-v.incoming.Queue
	if ok {
		return e
	}
	return nil
}

func (v *EventDispatcher) sendOutgoingEvent(event engine.Event, idx int) {
	v.outgoings[idx].Queue <- event
}

func (v *EventDispatcher) SetIncoming(queue *engine.EventQueue) {
	v.incoming = queue
}

func (v *EventDispatcher) SetOutgoings(queues []*engine.EventQueue) {
	v.outgoings = queues
}
