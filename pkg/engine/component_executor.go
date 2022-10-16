package engine

/**
 * The base class for executors of source and operator.
 */
type ComponentExecutor interface {
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
	parallelism       int
	instanceExecutors []InstanceExecutor
}

func (c *ComponentExecutorImpl) GetInstanceExecutors() []InstanceExecutor {
	return c.instanceExecutors
}

func (c *ComponentExecutorImpl) SetIncomingQueues(queues []*EventQueue) {
	for i := range queues {
		c.instanceExecutors[i].SetIncomingQueue(queues[i])
	}
}

func (c *ComponentExecutorImpl) SetOutgoingQueues(queue *EventQueue) {
	for i := range c.instanceExecutors {
		c.instanceExecutors[i].SetIncomingQueue(queue)
	}
}

func (c *ComponentExecutorImpl) Start() {
	panic("Need specific implementation")
}