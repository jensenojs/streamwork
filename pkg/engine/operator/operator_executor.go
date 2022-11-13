package operator

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/transport/strategy"

	"github.com/huandu/go-clone"
)

func NewOperatorExecutor(op engine.Operator) *OperatorExecutor {
	oe := new(OperatorExecutor)
	oe.Name = op.GetName()
	oe.Parallelism = op.GetParallelism()
	s := op.GetGroupingStrategy()
	if s == nil {
		oe.gs = strategy.NewShuffleGrouping()
	} else {
		oe.gs = s
	}

	oe.InstanceExecutors = make([]engine.InstanceExecutor, oe.Parallelism)
	for i := range oe.InstanceExecutors {
		// need clone new operator but not use the same one
		// otherwise parallelism will become meaningless
		c := clone.Clone(op).(engine.Operator)
		oe.InstanceExecutors[i] = NewOperatorExecutorInstance(i, c)
	}
	return oe
}

func (o *OperatorExecutor) GetParallelism() int {
	if o.Parallelism <= 0 || o.Parallelism > 10 {
		panic("An inappropriate number of concurrent requests")
	}
	return o.Parallelism
}

func (o *OperatorExecutor) GetGroupingStrategy() engine.GroupStrategy {
	return o.gs
}

func (o *OperatorExecutor) SetGroupingStrategy(gs engine.GroupStrategy) {
	o.gs = gs
}
