package operator

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/transport/strategy"

	"github.com/huandu/go-clone"
)

func NewOperatorExecutor(op engine.Operator) *OperatorExecutor {
	oe := &OperatorExecutor{
		gs: strategy.NewShuffleGrouping(), // default group strategy is round robin
	}
	oe.Parallelism = op.GetParallelism()
	oe.InstanceExecutors = make([]engine.InstanceExecutor, oe.Parallelism)
	for i := range oe.InstanceExecutors {
		// need clone new operator but not use the same one
		// otherwise parallelism will become meaningless
		c := clone.Clone(op).(engine.Operator)
		oe.InstanceExecutors[i] = NewOperatorExecutorInstance(i, c)
	}
	return oe
}

func (o *OperatorExecutor) GetGroupingStrategy() strategy.GroupStrategy {
	return o.gs
}

func (o *OperatorExecutor) SetGroupingStrategy(gs strategy.GroupStrategy) {
	o.gs = gs
}
