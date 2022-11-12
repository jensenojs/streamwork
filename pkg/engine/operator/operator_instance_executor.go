package operator

import (
	"fmt"
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/component"
)

func NewOperatorExecutorInstance(Id int, op engine.Operator) *OperatorInstanceExecutor {
	var opi = new(OperatorInstanceExecutor)
	opi.InstanceId = Id
	opi.operator = op
	opi.EventCollector = component.NewEventCollector()
	opi.OutgoingMap = make(map[engine.Channel][]engine.EventQueue)

	opi.SetRunOnce(opi.RunOnce)
	return opi
}

func (o *OperatorInstanceExecutor) RunOnce() bool {
	// read input
	event := o.TakeIncomingEvent()
	if event == nil {
		panic("receive nil event")
	}

	// apply operatorion
	fmt.Printf("%s:(%d) --> \n", o.operator.GetName(), o.InstanceId)
	o.operator.Apply(event, o.EventCollector)

	// emit out : should work.?
	o.SendOutgoingEvent()

	// clean up event that executed
	o.EventCollector.Clear()

	return true
}
