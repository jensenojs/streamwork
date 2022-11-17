package windowstrategy

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/transport"
	"time"
)

type SlidingTimeWindow struct {
	le time.Duration // Length of each window in milli seconds.
	in time.Duration // Interval between two adjacent windows. For fixed windows, interval should be equivalent to length.
	wa time.Duration // Extra wait time before window closing. A window should be closed when the current *processing time* is greater than the window end time (event time based) plus an extra watermark time.

	EventWindows map[time.Time]engine.EventWindow // Keep track of multiple event windows. Because of sliding and watermark, there could be multiple event windows existing in a strategy.
}

func NewSlidingTimeWindow(lengthMillis, intervalMillis, watermarkMillis time.Duration) *SlidingTimeWindow {
	return &SlidingTimeWindow{
		le:           lengthMillis,
		in:           intervalMillis,
		wa:           watermarkMillis,
		EventWindows: make(map[time.Time]engine.EventWindow),
	}
}

func (w *SlidingTimeWindow) IsLateEvent(eventTime, processTime time.Time) bool {
	return processTime.After(eventTime.Add(w.le).Add(w.wa))
}

func (w *SlidingTimeWindow) Add(e engine.Event, pt time.Time) {
	if te, ok := e.(engine.TimeEvent); !ok {
		panic("Timed events are required by time based WindowingStrategy")
	} else {
		et := te.GetTime()
		if !w.IsLateEvent(et, pt) {
			// Add event to all the windows cover its time
			mostRecentStart := et
			start := mostRecentStart
			for {
				if et.After(start.Add(w.le)) || et.Equal(start.Add(w.le)) {
					break
				}
				ew, ok := w.EventWindows[start]
				if !ok {
					w.EventWindows[start] = transport.NewEventWindow(start, start.Add(w.le))
				}
				ew.Add(e)
				start.Add(-w.in)
			}
		}
	}
}

// Get the event windows that are ready to be processed. It is based on the current processing time.
func (w *SlidingTimeWindow) GetEventWindows(processTime time.Time) []engine.EventWindow {
	// Return that windows that are ready to be processed. Typically there should be zero or one window.
	toProcess := make([]engine.EventWindow, 0)
	for startTime, v := range w.EventWindows {
		if processTime.After(startTime.Add(w.le).Add(w.wa)) {
			toProcess = append(toProcess, v)
		}
	}

	// clean up
	for _, k := range toProcess {
		delete(w.EventWindows, k.GetStartTime())
	}

	return toProcess
}
