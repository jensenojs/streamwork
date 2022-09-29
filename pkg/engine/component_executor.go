package engine

import "streamwork/pkg/api"

/**
 * The base class for executors of source and operator.
 */
type ComponentExecutor interface {
	GetComponent() api.Component

	SetIncomingQueue(i EventQueue)

	SetOutgoingQueue(i EventQueue)
}

/**
 * 	used to Inherited by operator_executor and source_executor
 */
type componentExecutor struct {
	component      api.Component
	fn             func()
	eventCollector []api.Event // accept events from user logic
	incomingQueue  EventQueue  // for upstream processes
	outgoingQueue  EventQueue  // for downstream processes
}

func NewComponentExecutor(component api.Component) *componentExecutor {
	return &componentExecutor{
		component: component,
	}
}

// =================================================================
// implement for ComponentExecutor
func (c *componentExecutor) GetComponent() api.Component {
	return c.component
}

func (c *componentExecutor) SetIncomingQueue(i EventQueue) {
	c.incomingQueue = i
}

func (c *componentExecutor) SetOutgoingQueue(o EventQueue) {
	c.outgoingQueue = o
}

// =================================================================
// implement for Process
func (c *componentExecutor) Process() {
	c.fn = func() {
		go func() {
			for {
				c.runOnce()
			}
		}()
	}
}

func (c *componentExecutor) Start() {
	go c.fn()
}

func (c *componentExecutor) runOnce() bool {
	panic("Need to implement runOnce")
}

// =================================================================
// some addtional helper functions
func (c *componentExecutor) takeIncomingEvent() api.Event {
	e, ok := <-c.incomingQueue.queue
	if ok {
		return e
	}
	return nil
}

func (c *componentExecutor) sendOutgoingEvent(event api.Event) {
	c.outgoingQueue.queue <- event
}
