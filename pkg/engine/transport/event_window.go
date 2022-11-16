package transport

import (
	"streamwork/pkg/engine"
	"time"
)

func NewEventWindow(s, e time.Time) *EventWindow {
	return &EventWindow{
		start: s,
		end:   e,
	}
}

// Get the start timestamp of the window. The time is inclusive.
func (e *EventWindow) GetStartTime() time.Time {
	return e.start
}

// Get the end timestamp of the window. The time is exclusive.
func (e *EventWindow) GetEndTime() time.Time {
	return e.end
}

func (e *EventWindow) Add(ev engine.Event) {
	e.List = append(e.List, ev)
}

func (e *EventWindow) GetEvents() []engine.Event {
	return e.List
}
