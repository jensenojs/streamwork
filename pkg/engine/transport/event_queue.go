package transport

import "streamwork/pkg/engine"

func NewEventQueue(arg ...int) *EventQueue {
	var size int
	if (len(arg) == 0) {
		size = 64
	} else if (len(arg) == 1){
		size = arg[0]
	} else {
		panic("Too many arguments for event queue init")
	}

	return &EventQueue{
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
