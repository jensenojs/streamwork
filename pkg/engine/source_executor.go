package engine

import "streamwork/pkg/api"

/**
 * The executor for source components. When the executor is started,
 * a new thread is created to call the getEvents() function of
 * the source component repeatedly.
 *
 * Used to inherited by specific operator
 */
type SourceExecutor struct {
	ComponentExecutorImpl
}

// only used in job_starter
func newSourceExecutor(s api.Source) *SourceExecutor {
	// needs to set or read fields by func
	return &SourceExecutor{}
}

func (s *SourceExecutor) GetEvents([]api.Event) error {
	panic("Need to be implemented by specific source")
}

func (s *SourceExecutor) runOnce() bool {
	// generate events
	if s.GetEvents(s.eventCollector) != nil {
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
