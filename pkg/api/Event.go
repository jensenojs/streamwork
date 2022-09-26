package api

/**
 * This is the base class for all the event classes.
 * Users should extend this class to implement all their own event classes.
 */
type Event interface {
	// Get data stored in the event.
	GetData() any
}
