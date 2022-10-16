package api

/**
 * This Operator class is the base class for all user defined operators.
 */
type Operator interface {
	Component

	/**
	 * Apply logic to the incoming event and generate results.
	 * The function is abstract and needs to be implemented by users.
	 * @param event The incoming event
	 * @param eventCollector The outgoing event collector
	 */
	Apply(Event, *[]Event) error

	// set up instance
	SetupInstance(instanceId int)

	// get the grouping strategy
	GetGroupingStrategy() GroupStrategy
}
