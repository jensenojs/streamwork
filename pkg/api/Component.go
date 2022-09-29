package api

/**
 * The base class for all components, including Source and Operator.
 */
type Component interface {
	GetName() string

	// Get the outgoing event stream of this component. The stream is used to connect the downstream components.
	GetOutgoingStream() *Stream
}

type Componentimpl struct {
	name string
	// The stream object is used to connect the downstream operators.
	outgoingStream *Stream
}

func (c *Componentimpl) GetName() string {
	return c.name
}


func (c *Componentimpl) GetOutgoingStream() *Stream {
	return c.outgoingStream
}