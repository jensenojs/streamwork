package api

/**
 * The base class for all components, including Source and Operator.
 */
type Component struct {
	name string
	// The stream object is used to connect the downstream operators.
	outgoingStream Stream
}
 
func NewComponent(name string) *Component {
	return &Component{
		name: name,
		outgoingStream: NewStream(),
	}
}


func (c *Component) GetName() string {
	return c.name
}

/**
 * Get the outgoing event stream of this component. The stream is used to connect
 * the downstream components.
 */
func (c *Component) GetOutgoingStream() Stream {
	return c.outgoingStream
}