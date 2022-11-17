package windowstrategy

import (
	"streamwork/pkg/engine"
	"time"
)

type FixedTimeWindow struct {
	SlidingTimeWindow
}

func NewFixedTimeWindow(intervalMillis, watermarkMillis time.Duration) *FixedTimeWindow {
	f := new(FixedTimeWindow)

	f.le = intervalMillis
	f.in = intervalMillis
	f.wa = watermarkMillis
	f.EventWindows = make(map[time.Time]engine.EventWindow)
	return f
}
