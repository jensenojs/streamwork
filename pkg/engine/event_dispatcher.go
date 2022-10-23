package engine

import "streamwork/pkg/api"

/**
 * EventDispatcher is responsible for transporting events from
 * the incoming queue to the outgoing queues with a grouping strategy.
 */
type EventDispatcher struct {
	fnWrapper          func() // wrapper function for fn, no need for fn
	downStreamExecutor *OperatorExecutor
	incomingQueue      *EventQueue
	outgoingQueues     []*EventQueue
}

func NewEventDispatcher(downStreamExecutor *OperatorExecutor) *EventDispatcher {
	return &EventDispatcher{
		downStreamExecutor: downStreamExecutor,
	}
}

// ============================================================
// implementation of Process
func (v *EventDispatcher) newProcess() {
	v.fnWrapper = func() {
		for {
			if ok := v.runOnce(); ok != true {
				break
			}
		}
	}
}

func (v *EventDispatcher) Start() {
	go v.fnWrapper()
}

func (v *EventDispatcher) runOnce() bool {
	e := v.takeIncomingEvent()
	s := v.downStreamExecutor.GetGroupingStrategy()
	idx := s.GetInstance(e, len(v.outgoingQueues))
	v.sendOutgoingEvent(e, idx)
	return true
}

// ============================================================
// init and use for incoming/outgoing queues
func (v *EventDispatcher) takeIncomingEvent() api.Event {
	e, ok := <-v.incomingQueue.queue
	if ok {
		return e
	}
	return nil
}

func (v *EventDispatcher) sendOutgoingEvent(event api.Event, idx int) {
	v.outgoingQueues[idx].queue <- event
}

func (v *EventDispatcher) SetIncomingQueue(queue *EventQueue) {
	v.incomingQueue = queue
}

func (v *EventDispatcher) SetOutgoingQueues(queues []*EventQueue) {
	v.outgoingQueues = queues
}
