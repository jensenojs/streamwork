package source

import (
	"streamwork/pkg/engine"
)

// =================================================================
// implement for Source

func (s *Source) GetEvents([]byte, int, engine.EventCollector) {
	panic("need to implement GetEvents")
}
