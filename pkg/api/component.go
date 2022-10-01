package api

/**
 * The base class for all components, including Source and Operator.
 */
type Component interface {
	SetName(name string)

	GetName() string

	SetOutgoingStream()

	// Get the outgoing event stream of this component. The stream is used to connect the downstream components.
	GetOutgoingStream() *Stream
}