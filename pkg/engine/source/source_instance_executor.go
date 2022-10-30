package source

import "streamwork/pkg/engine"

func NewSourceExecutorInstance(id int, so engine.Source) *SourceInstanceExecutor {
	var soi = new(SourceInstanceExecutor)
	soi.InstanceId = id
	soi.source = so
	soi.source.SetupInstance(id)
	soi.SetRunOnce(soi.RunOnce)
	return soi
}

func (s *SourceInstanceExecutor) RunOnce() bool {
	// generate events
	s.source.GetEvents(&s.EventCollector)

	// emit out
	for _, e := range s.EventCollector {
		s.SendOutgoingEvent(e)
	}

	// clean up event that executed
	s.EventCollector = nil

	return true
}
