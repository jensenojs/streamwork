package source

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/component"
)

func NewSourceExecutorInstance(id int, so engine.Source) *SourceInstanceExecutor {
	var soi = new(SourceInstanceExecutor)
	soi.InstanceId = id
	soi.source = so
	soi.source.SetupInstance(id) // really need this?
	soi.SetRunOnce(soi.RunOnce)
	soi.EventCollector = component.NewEventCollector()
	return soi
}

func (s *SourceInstanceExecutor) RunOnce() bool {
	// generate events
	s.source.GetEvents(s.EventCollector)

	// emit out
	s.SendOutgoingEvent()

	// clean up event that executed
	s.EventCollector.Clear()

	return true
}
