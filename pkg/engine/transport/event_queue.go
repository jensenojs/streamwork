package transport

import "streamwork/pkg/engine"

func NewEventQueue(size int) *engine.EventQueue {
	return &engine.EventQueue{
		Queue: make(chan engine.Event, size),
	}
}
