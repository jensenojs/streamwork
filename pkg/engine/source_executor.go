package engine

import "streamwork/pkg/api"

/**
 * The executor for source components. When the executor is started,
 * a new thread is created to call the getEvents() function of
 * the source component repeatedly.
 *
 * Used to inherited by specific operator
 */
type SourceExecutor struct {
	ComponentExecutorImpl
	so api.Source // specific source, used to execute GetEvent
}

func newSourceExecutor(so api.Source) *SourceExecutor {
	se := &SourceExecutor{
		so: so,
	}
	se.parallelism = so.GetParallelism()
	se.instanceExecutors = make([]InstanceExecutor, se.parallelism)
	for i := range se.instanceExecutors {
		se.instanceExecutors[i] = newSourceExecutorInstance(i, so)
	}
	return se
}

func (s *SourceExecutor) Start() {
	if s.instanceExecutors == nil {
		panic("Should not be nil")
	}
	for i := range s.instanceExecutors {
		s.instanceExecutors[i].Start()
	}
}

func (s *SourceExecutor) SetIncomingQueues(queues []*EventQueue) {
	panic("No incoming queue is allowed for source executor")
}
