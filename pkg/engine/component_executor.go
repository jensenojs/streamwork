package engine

import (
	"streamwork/pkg/api"
)

/**
 * The base class for executors of source and operator.
 */
type ComponentExecutor interface {
	api.Component

	// Get the instance executors of this component executor.
	GetInstanceExecutors() []InstanceExecutor

	SetIncomingQueues(queues []*EventQueue)

	SetOutgoingQueue(queue *EventQueue)

	Start()
}

/**
 * 	used to Inherited by OperatorExecutor and SourceExecutor
 */
type ComponentExecutorImpl struct {
	name              string
	stream            *api.Stream // connect to next component
	parallelism       int
	instanceExecutors []InstanceExecutor
}

// =================================================================
// Component
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

func (c *ComponentExecutorImpl) SetParallelism(parallelism int) {
	if parallelism < 0 || parallelism > 10 {
		panic("Inappropriate parallelism number")
	}
	c.parallelism = parallelism
}

func (c *ComponentExecutorImpl) GetParallelism() int {
	return c.parallelism
}

// helper function to init a instance executor combine SetName, SetOutgoingStream and SetParallelism
func (c *ComponentExecutorImpl) Init(name string, parallelism int) {
	c.SetName(name)
	c.SetOutgoingStream()
	c.SetParallelism(parallelism)
}

// =================================================================
// implement for ComponentExecutor
func (c *ComponentExecutorImpl) GetInstanceExecutors() []InstanceExecutor {
	return c.instanceExecutors
}

func (c *ComponentExecutorImpl) SetIncomingQueues(queues []*EventQueue) {
	for i := range queues {
		c.instanceExecutors[i].SetIncomingQueue(queues[i])
	}
}

func (c *ComponentExecutorImpl) SetOutgoingQueue(queue *EventQueue) {
	for i := range c.instanceExecutors {
		c.instanceExecutors[i].SetIncomingQueue(queue)
	}
}

func (c *ComponentExecutorImpl) Start() {
	panic("Need specific implementation")
}
