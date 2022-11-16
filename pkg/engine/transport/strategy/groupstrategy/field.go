package groupstrategy

import (
	"streamwork/pkg/engine"
)

// FieldGrouping the event are routed to downstream
// instances by its value. This implementation has many limitations
// as it only supports mapping events by string.
type FieldGrouping struct {
	cnt          int
	Map          map[string]int
	CustomGetKey func(engine.Event) string
}

// Get key from an event. Child class can override this function to calculate key in different ways.
// For example, calculate the key from some specific fields.
func (f *FieldGrouping) GetKey(event engine.Event) string {
	if f.CustomGetKey == nil {
		panic("please implement custom GetKey function")
	}
	return f.CustomGetKey(event)
}

// Get target instance id from an event and component parallelism.
func (f *FieldGrouping) GetInstance(event engine.Event, parallelism int) int {
	s := f.GetKey(event)
	val, ok := f.Map[s]
	if !ok {
		f.Map[s] = f.cnt
		f.cnt++
		return f.Map[s] % parallelism
	}
	return val % parallelism
}
