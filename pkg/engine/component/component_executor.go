package component

import (
	"streamwork/pkg/engine"
)

// =================================================================
// implement for Component
func (c *ComponentExecutorImpl) SetName(name string) { // Use InitNameAndStream to instead
	c.Name = name
}

func (c *ComponentExecutorImpl) GetName() string {
	return c.Name
}

func (c *ComponentExecutorImpl) SetOutgoingStream() { // Use InitNameAndStream to instead
	if c.Stream == nil {
		c.Stream = engine.NewStream()
	}
}

func (c *ComponentExecutorImpl) GetOutgoingStream() *engine.Stream {
	return c.Stream
}

func (c *ComponentExecutorImpl) SetParallelism(parallelism int) {
	if parallelism < 0 || parallelism > 10 {
		panic("Inappropriate parallelism number")
	}
	c.Parallelism = parallelism
}

func (c *ComponentExecutorImpl) GetParallelism() int {
	return c.Parallelism
}

// helper function to init a instance executor combine SetName, SetOutgoingStream and SetParallelism
func (c *ComponentExecutorImpl) Init(name string, parallelism int) {
	c.SetName(name)
	c.SetOutgoingStream()
	c.SetParallelism(parallelism)
}

// =================================================================
// implement for ComponentExecutor
func (c *ComponentExecutorImpl) GetInstanceExecutors() []engine.InstanceExecutor {
	return c.InstanceExecutors
}

func (c *ComponentExecutorImpl) SetIncomings(queues []*engine.EventQueue) {
	for i := range queues {
		c.InstanceExecutors[i].SetIncoming(queues[i])
	}
}

func (c *ComponentExecutorImpl) SetOutgoing(queue *engine.EventQueue) {
	for i := range c.InstanceExecutors {
		c.InstanceExecutors[i].SetOutgoing(queue)
	}
}

// =================================================================
// implement for process
func (c *ComponentExecutorImpl) NewProcess() {
	if c.InstanceExecutors == nil {
		panic("Should not be nil")
	}
	for i := range c.InstanceExecutors {
		c.InstanceExecutors[i].NewProcess()
	}
}

func (c *ComponentExecutorImpl) Start() {
	if c.InstanceExecutors == nil {
		panic("Should not be nil")
	}
	for i := range c.InstanceExecutors {
		c.InstanceExecutors[i].Start()
	}
}

func (c *ComponentExecutorImpl) RunOnce() bool {
	panic("Need specific implementation")
}
