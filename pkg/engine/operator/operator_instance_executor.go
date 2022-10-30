package operator

import "streamwork/pkg/engine"

func NewOperatorExecutorInstance(Id int, op engine.Operator) *OperatorInstanceExecutor {
	var opi = new(OperatorInstanceExecutor)
	opi.InstanceId = Id
	opi.operator = op
	opi.operator.SetupInstance(Id)
	opi.SetRunOnce(opi.RunOnce)
	return opi
}

func (o *OperatorInstanceExecutor) RunOnce() bool {
	// read input
	event := o.TakeIncomingEvent()
	if event == nil {
		return false
	}

	// apply operatorion
	o.operator.Apply(event, &o.EventCollector)

	// emit out : should work.?
	for _, e := range o.EventCollector {
		o.SendOutgoingEvent(e)
	}

	// clean up event that executed
	o.EventCollector = nil

	return true
}
