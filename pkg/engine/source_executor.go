package engine

import "streamwork/pkg/api"

type SourceExecutor struct {
	componentExecutor
	s api.Source
}

func NewSourceExecutor(s api.Source) *SourceExecutor {
	return &SourceExecutor{
		s: s,
	}
}

func (s *SourceExecutor) runOnce() bool {
	// generate events
	if s.s.GetEvents(s.eventCollector) != nil {
		return false
	}

	// emit out
	for _, e := range s.eventCollector {
		s.sendOutgoingEvent(e)
	}
	return true
}

func (s *SourceExecutor) SetIncomingQueue(i *EventQueue) {
	panic("No incoming queue is allowed for source executor")
}
