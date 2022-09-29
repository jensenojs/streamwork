package engine

import "streamwork/pkg/api"

/**
 * The executor for operator components. When the executor is started, a new thread
 * is created to call the apply() function of the operator component repeatedly.
 */
type OperatorExecutor struct {
	componentExecutor
	op api.Operator
}

func NewOperatorExecutor(o api.Operator) *OperatorExecutor {
	return &OperatorExecutor{
		op: o,
	}
}

/* Run process once.
 * @return true if the thread should continue; false if the thread should exist.
 */
func (o *OperatorExecutor) runOnce() bool {
	// read input
	event := o.takeIncomingEvent()
	if event == nil {
		return false
	}

	// apply operatorion
	o.op.Apply(event, o.eventCollector)

	// emit out : should work.?
	for _, e := range o.eventCollector {
		o.sendOutgoingEvent(e)
	}
	o.eventCollector = nil

	return true
}
