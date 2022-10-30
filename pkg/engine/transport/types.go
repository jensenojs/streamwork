package transport

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/operator"
)

// EventDispatcher is responsible for transporting events from
// the incoming queue to the outgoing queues with a grouping strategy.
type EventDispatcher struct {
	fnWrapper          func() // wrapper function for fn, no need for fn
	downStreamExecutor *operator.OperatorExecutor
	incoming           engine.EventQueue
	outgoings          []engine.EventQueue
}

/**
 * This is the class for intemediate event queues between processes.
 */
type EventQueue struct {
	Queue chan engine.Event
}

