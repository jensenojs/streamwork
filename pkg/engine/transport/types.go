package transport

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/operator"
	"time"
)

// EventDispatcher is responsible for transporting events from
// the incoming queue to the outgoing queues with a grouping strategy.
type EventDispatcher struct {
	fnWrapper          func() // wrapper function for RunOnce. It has nothing to do with user configuration logic, so there is no need for fn
	downStreamExecutor *operator.OperatorExecutor
	incoming           engine.EventQueue
	outgoings          []engine.EventQueue
}

// EventQueue is responsible for intemediate event queues between processes.
type EventQueue struct {
	Queue chan engine.Event
}

// EventWindow is responsible for collect event into a window
type EventWindow struct {
	List  []engine.Event
	start time.Time
	end   time.Time
}
