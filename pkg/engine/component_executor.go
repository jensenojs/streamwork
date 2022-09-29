package engine

import "streamwork/pkg/api"

/**
 * 	used to Inherited by a specific operator
 */
type componentExecutor struct {
	process
	ComponentExecutor

	component      api.Component
	eventCollector []api.Event // accept events from user logic
	incomingQueue  EventQueue  // for upstream processes
	outgoingQueue  EventQueue  // for downstream processes
}

func (c *componentExecutor) SetIncomingQueue(i EventQueue) {
	c.incomingQueue = i
}

func (c *componentExecutor) SetOutgoingQueue(o EventQueue) {
	c.outgoingQueue = o
}
