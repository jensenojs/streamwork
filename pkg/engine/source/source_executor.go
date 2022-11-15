package source

import (
	"streamwork/pkg/engine"

	"github.com/huandu/go-clone"
)

var DEFAULT_USED bool

func NewSourceExecutor(so engine.Source) *SourceExecutor {
	se := new(SourceExecutor)
	se.Name = so.GetName()
	se.Parallelism = so.GetParallelism()
	se.InstanceExecutors = make([]engine.InstanceExecutor, se.Parallelism)
	if !DEFAULT_USED {
		se.PortBase = ConnPort
		DEFAULT_USED = true
	} else {
		se.PortBase = ConnPort2
	}

	for i := range se.InstanceExecutors {
		// need clone new operator but not use the same one
		// otherwise parallelism will become meaningless
		c := clone.Clone(so).(engine.Source)
		se.InstanceExecutors[i] = NewSourceExecutorInstance(se.PortBase, i, c)
	}
	return se
}

func (s *SourceExecutor) SetIncomings(queues []engine.EventQueue) {
	panic("No incoming queue is allowed for source executor")
}
