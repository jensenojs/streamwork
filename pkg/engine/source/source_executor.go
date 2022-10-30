package source

import (
	"streamwork/pkg/engine"

	"github.com/huandu/go-clone"
)

func NewSourceExecutor(so engine.Source) *SourceExecutor {
	se := &SourceExecutor{
		so: so,
	}
	se.Parallelism = so.GetParallelism()
	se.InstanceExecutors = make([]engine.InstanceExecutor, se.Parallelism)
	for i := range se.InstanceExecutors {
		// need clone new operator but not use the same one
		// otherwise parallelism will become meaningless
		c := clone.Clone(so).(engine.Source)
		se.InstanceExecutors[i] = NewSourceExecutorInstance(i, c)
	}
	return se
}

func (s *SourceExecutor) SetIncomings(queues []engine.EventQueue) {
	panic("No incoming queue is allowed for source executor")
}
