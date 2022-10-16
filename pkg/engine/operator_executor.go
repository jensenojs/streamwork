package engine

import "streamwork/pkg/api"

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
		oe.instanceExecutors[i] = newOperatorExecutorInstance(i, op)
	}
	return oe
}

func (o *OperatorExecutor) Start() {
	if o.instanceExecutors == nil {
		panic("Should not be nil")
	}
	for i := range o.instanceExecutors {
		o.instanceExecutors[i].Start()
	}
}

func (o *OperatorExecutor) GetGroupingStrategy() api.GroupStrategy {
	return o.gs
}

func (o *OperatorExecutor) SetGroupingStrategy(gs api.GroupStrategy) {
	o.gs = gs
}
