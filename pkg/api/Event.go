package api

/**
 * This is the base class for all the event classes.
 * Users should extend this class to implement all their own event classes.
 */
type Event struct {
}

 
// Get data stored in the event.
func (e *Event) GetData() any{
	panic("Need to implement GetData")
}