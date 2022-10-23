package strategy

import "streamwork/pkg/engine"

// ShuffleGrouping the events are routed to downstream
// instances relatively evenly. This implementation is round robin
// based instead of random number based because it is simpler and
// deterministic.
type ShuffleGrouping struct {
	cnt int
}

func NewShuffleGrouping() *ShuffleGrouping {
	return &ShuffleGrouping{
		cnt: 0,
	}
}

func (s *ShuffleGrouping) GetInstance(event engine.Event, parallelism int) int {
	if s.cnt >= parallelism {
		s.cnt = 0
	}
	s.cnt++
	return s.cnt - 1
}
