package component

import (
	"streamwork/pkg/engine"
)

// =================================================================
// implement for ComponentExecutor

func (c *ComponentExecutorImpl) GetInstanceExecutors() []engine.InstanceExecutor {
	return c.InstanceExecutors
}

func (c *ComponentExecutorImpl) SetIncomings(queues []engine.EventQueue) {
	for i := range queues {
		c.InstanceExecutors[i].SetIncoming(queues[i])
	}
}

func (c *ComponentExecutorImpl) AddOutgoing(ch engine.Channel, qu engine.EventQueue) {
	for i := range c.InstanceExecutors {
		c.InstanceExecutors[i].AddOutgoing(ch, qu)
	}
}

func (c *ComponentExecutorImpl) RegisterChannel(ch engine.Channel) {
	for i := range c.InstanceExecutors {
		c.InstanceExecutors[i].RegisterChannel(ch)
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
