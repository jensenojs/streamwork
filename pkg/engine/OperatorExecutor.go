package engine

import "streamwork/pkg/api"
/**
 * The executor for operator components. When the executor is started, a new thread
 * is created to call the apply() function of the operator component repeatedly.
 */
type OperatorExecutor struct {
	componentExecutor
	operator api.Operator
}

func NewOperatorExecutor(op api.Operator) *OperatorExecutor {
	return &OperatorExecutor{
		operator: op,
	}
}

func (e *OperatorExecutor) runOnce() bool {
	return false
}