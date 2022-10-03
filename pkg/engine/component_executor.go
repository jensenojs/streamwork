package engine

import "streamwork/pkg/api"

/**
 * The base class for executors of source and operator.
 */
type ComponentExecutor interface {
	api.Component

	SetIncomingQueue(i *EventQueue)

	SetOutgoingQueue(i *EventQueue)
}

/**
 * 	used to Inherited by OperatorExecutor and SourceExecutor
 */
type ComponentExecutorImpl struct {
	name           string
	fn             func()
	eventCollector []api.Event // accept events from user logic
	stream         *api.Stream
	incomingQueue  *EventQueue // for upstream processes
	outgoingQueue  *EventQueue // for downstream processes
}

// =================================================================
// implement for Component
func (c *ComponentExecutorImpl) SetName(name string) { // Use InitNameAndStream to instead
	c.name = name
}

func (c *ComponentExecutorImpl) GetName() string {
	return c.name
}

func (c *ComponentExecutorImpl) SetOutgoingStream() { // Use InitNameAndStream to instead
	if c.stream == nil {
		c.stream = api.NewStream()
	}
}

func (c *ComponentExecutorImpl) GetOutgoingStream() *api.Stream {
	return c.stream
}

// helper function to init a component executor, trying to not use SetName or SetOutgoingStream
func (c *ComponentExecutorImpl) InitNameAndStream(name string) {
	c.SetName(name)
	c.SetOutgoingStream()
}

// =================================================================
// implement for ComponentExecutor
func (c *ComponentExecutorImpl) SetIncomingQueue(i *EventQueue) {
	c.incomingQueue = i
}

func (c *ComponentExecutorImpl) SetOutgoingQueue(o *EventQueue) {
	c.outgoingQueue = o
}

// helper functions to receive/send events
func (c *ComponentExecutorImpl) takeIncomingEvent() api.Event {
	e, ok := <-c.incomingQueue.queue
	if ok {
		return e
	}
	return nil
}

func (c *ComponentExecutorImpl) sendOutgoingEvent(event api.Event) {
	c.outgoingQueue.queue <- event
}

// =================================================================
// implement for Process
func (c *ComponentExecutorImpl) Process() {
	c.fn = func() {
		go func() {
			for {
				c.runOnce()
			}
		}()
	}
}

func (c *ComponentExecutorImpl) Start() {
	go c.fn()
}

func (c *ComponentExecutorImpl) runOnce() bool {
	panic("Need to implement runOnce")
}
