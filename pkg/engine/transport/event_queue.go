package transport

import "streamwork/pkg/engine"

func NewEventQueue(size int) EventQueue {
	return EventQueue{
		Queue: make(chan engine.Event, size),
	}
}

func (q *EventQueue) Take() engine.Event {
	e, ok := <-q.Queue
	if ok {
		return e
	}
	return nil
}

func (q *EventQueue) Send(e engine.Event) {
	q.Queue <- e
}
