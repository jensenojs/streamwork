package api

/**
 * This Source class is the base class for all user defined sources.
 */
type Source struct {
	Component
}

/**
 * Accept events from external into the system.
 * The function is abstract and needs to be implemented by users.
 * @param eventCollector The outgoing event collector
 */
func GetEvents(eventCollector []Event) (err error) {
	panic("Need to implement GetEvents")
}
