package operator

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/component"
)

func NewOperatorExecutorInstance(Id int, op engine.Operator) *OperatorInstanceExecutor {
	var opi = new(OperatorInstanceExecutor)
	opi.InstanceId = Id
	opi.operator = op
	opi.operator.SetupInstance(Id) // really need this?
	opi.SetRunOnce(opi.RunOnce)
	opi.EventCollector = component.NewEventCollector()
	opi.OutgoingMap = make(map[engine.Channel][]engine.EventQueue)
	return opi
}

func (o *OperatorInstanceExecutor) RunOnce() bool {
	// read input
	event := o.TakeIncomingEvent()
	if event == nil {
		panic("receive nil event")
	}

	// apply operatorion
	o.operator.Apply(event, o.EventCollector)

	// emit out : should work.?
	o.SendOutgoingEvent()

	// clean up event that executed
	o.EventCollector.Clear()

	return true
}
