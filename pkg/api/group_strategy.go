package api

/**
 * Get target instance id from an event and component parallelism.
 * Note that in this implementation, only one instance is selected.
 * This can be easily extended if needed.
 */
type GroupStrategy interface {
	// the event object to route to the component
	// the parallelism of the component
	GetInstance(event Event, parallelism int) int
}