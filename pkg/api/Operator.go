package api

/**
 * This Operator class is the base class for all user defined operators.
 */
type Operator struct {
	Component
}

/**
 * Apply logic to the incoming event and generate results.
 * The function is abstract and needs to be implemented by users.
 * @param event The incoming event
 * @param eventCollector The outgoing event collector
 */
func Apply(event Event, eventCollector []Event) error {
	panic("Need to implement Apply")
}
