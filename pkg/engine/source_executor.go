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
	source api.Source
}

func newSourceExecutor(s api.Source) *SourceExecutor {
	// needs to set or read fields by func
	se := &SourceExecutor{
		source: s,
	}
	// se.setRunOnce(se.runOnce)
	return se
}

func (s *SourceExecutor) GetEvents([]api.Event) {
	panic("Need to be implemented by specific source")
}

func (s *SourceExecutor) runOnce() bool {
	// get

	// generate events
	// s.source.GetEvents(&s.eventCollector)

	// emit out
	// for _, e := range s.eventCollector {
	// 	s.sendOutgoingEvent(e)
	// }

	// clean up event that executed
	// s.eventCollector = nil

	return true
}

func (s *SourceExecutor) SetIncomingQueue(i *EventQueue) {
	panic("No incoming queue is allowed for source executor")
}
