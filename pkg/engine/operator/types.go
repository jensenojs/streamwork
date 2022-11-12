package operator

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/component"
)

type Operator struct {
	component.Component
	Strategy engine.GroupStrategy
}

// The executor for operator components. When the executor is started,
// a new thread is created to call the apply() function of
// the operator component repeatedly.
//
// Used to inherited by specific operator
type OperatorExecutor struct {
	component.ComponentExecutorImpl
	gs engine.GroupStrategy // group strategy, different from origin implementation, place strategy in operatorExecutor but not operator
}

type OperatorInstanceExecutor struct {
	component.InstanceExecutorImpl
	operator engine.Operator
}
