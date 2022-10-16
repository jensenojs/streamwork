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
	op api.Operator // specific operator, used to execute apply
}

func newOperatorExecutor(op api.Operator) *OperatorExecutor {
	oe := &OperatorExecutor{
		op: op,
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
	return o.op.GetGroupingStrategy()
}
