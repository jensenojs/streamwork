package engine

import "streamwork/pkg/api"

/**
 * This is the class for intemediate event queues between processes.
 */
type EventQueue struct {
	queue chan api.Event
}

func NewEventQueue(size int) *EventQueue {
	return &EventQueue{
		queue: make(chan api.Event, size),
	}
}
