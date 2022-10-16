package api

/**
 * This Source class is the base class for all user defined sources.
 */
type Source interface {
	Component

	/**
	 * Accept events from external into the system.
	 * The function is abstract and needs to be implemented by users.
	 * @param eventCollector The outgoing event collector
	 */
	GetEvents(eventCollector *[]Event)

	// set up instance
	SetupInstance(instanceId int)
}
