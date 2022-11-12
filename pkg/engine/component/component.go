package component

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/stream"
)

// =================================================================
// implement for Component

func (c *Component) GetName() string {
	return c.Name
}

func (c *Component) GetOutgoingStream() engine.Stream {
	if c.Stream == nil {
		c.Stream = stream.NewStream()
	}
	return c.Stream
}

func (c *Component) GetParallelism() int {
	if c.Parallelism <= 0 || c.Parallelism > 10 {
		panic("An inappropriate number of concurrent requests")
	}
	return c.Parallelism
}
