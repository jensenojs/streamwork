package engine

import "streamwork/pkg/api"

type OperatorInstanceExecutor struct {
	InstanceExecutorImpl
	instanceId int
	operator   api.Operator
}

func newOperatorExecutorInstance(Id int, op api.Operator) *OperatorInstanceExecutor {
	var opi = new(OperatorInstanceExecutor)
	opi.instanceId = Id
	opi.operator = op
	opi.operator.SetupInstance(Id)
	opi.setRunOnce(opi.runOnce)
	return opi
}

/* Run process once.
 * @return true if the thread should continue; false if the thread should exist.
 */
func (o *OperatorInstanceExecutor) runOnce() bool {
	// read input
	event := o.takeIncomingEvent()
	if event == nil {
		return false
	}

	// apply operatorion
	o.operator.Apply(event, &o.eventCollector)

	// emit out : should work.?
	for _, e := range o.eventCollector {
		o.sendOutgoingEvent(e)
	}

	// clean up event that executed
	o.eventCollector = nil

	return true
}
