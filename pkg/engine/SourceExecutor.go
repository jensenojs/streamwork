package engine

import "streamwork/pkg/api"

type SourceExecutor struct {
	componentExecutor
	source api.Source
}

func NewSourceExecutor(s api.Source) *SourceExecutor {
	return &SourceExecutor{
		source: s,
	}
}

func (s *SourceExecutor) runOnce() bool {
	
	return false
}

func (s *SourceExecutor) SetIncomingQueue(i EventQueue) {
	panic("No incoming queue is allowed for source executor")
}