package engine

import (
	"streamwork/pkg/api"

	"github.com/huandu/go-clone"
)

/**
 * The executor for operator components. When the executor is started,
 * a new thread is created to call the apply() function of
 * the operator component repeatedly.
 *
 * Used to inherited by specific operator
 */
type OperatorExecutor struct {
	ComponentExecutorImpl
	op api.Operator      // specific operator, used to execute apply
	gs api.GroupStrategy // group strategy, different from origin implementation, place strategy in operatorExecutor but not operator
}

func newOperatorExecutor(op api.Operator) *OperatorExecutor {
	oe := &OperatorExecutor{
		op: op,
		gs: api.NewShuffleGrouping(), // default group strategy is round robin
	}
	oe.parallelism = op.GetParallelism()
	oe.instanceExecutors = make([]InstanceExecutor, oe.parallelism)
	for i := range oe.instanceExecutors {
		// need clone new operator but not use the same one
		// otherwise parallelism will become meaningless
		c := clone.Clone(op).(api.Operator)
		oe.instanceExecutors[i] = newOperatorExecutorInstance(i, c)
	}
	return oe
}

func (o *OperatorExecutor) GetGroupingStrategy() api.GroupStrategy {
	return o.gs
}

func (o *OperatorExecutor) SetGroupingStrategy(gs api.GroupStrategy) {
	o.gs = gs
}
