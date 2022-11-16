package windowstrategy

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/transport"
	"time"
)

type SlidingTimeWindow struct {
	Le time.Duration // Length of each window in milli seconds.
	In time.Duration // Interval between two adjacent windows. For fixed windows, interval should be equivalent to length.
	Wa time.Duration // Extra wait time before window closing. A window should be closed when the current *processing time* is greater than the window end time (event time based) plus an extra watermark time.

	eventWindows map[time.Time]engine.EventWindow // Keep track of multiple event windows. Because of sliding and watermark, there could be multiple event windows existing in a strategy.
}

func NewSlidingTimeWindow(lengthMillis, intervalMillis, watermarkMillis time.Duration) *SlidingTimeWindow {
	return &SlidingTimeWindow{
		Le:           lengthMillis,
		In:           intervalMillis,
		Wa:           watermarkMillis,
		eventWindows: make(map[time.Time]engine.EventWindow),
	}
}

func (w *SlidingTimeWindow) IsLateEvent(eventTime, processTime time.Time) bool {
	return processTime.After(eventTime.Add(w.Le).Add(w.Wa))
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
				if et.After(start.Add(w.Le)) || et.Equal(start.Add(w.Le)) {
					break
				}
				ew, ok := w.eventWindows[start]
				if !ok {
					w.eventWindows[start] = transport.NewEventWindow(start, start.Add(w.Le))
				}
				ew.Add(e)
				start.Add(-w.In)
			}
		}
	}
}

// Get the event windows that are ready to be processed. It is based on the current processing time.
func (w *SlidingTimeWindow) GetEventWindows(processTime time.Time) []engine.EventWindow {
	// Return that windows that are ready to be processed. Typically there should be zero or one window.
	toProcess := make([]engine.EventWindow, 0)
	for startTime, v := range w.eventWindows {
		if processTime.After(startTime.Add(w.Le).Add(w.Wa)) {
			toProcess = append(toProcess, v)
		}
	}

	// clean up
	for _, k := range toProcess {
		delete(w.eventWindows, k.GetStartTime())
	}

	return toProcess
}
