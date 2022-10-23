package engine

type Void struct{}

var Member Void

/**
 * The base class for all components, including Source and Operator.
 */
type Component interface {
	SetName(name string)

	GetName() string

	SetOutgoingStream()

	// Get the outgoing event stream of this component. The stream is used to connect the downstream components.
	GetOutgoingStream() *Stream

	// Get the parallelism (number of instances) of this component.
	GetParallelism() int
}

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
}

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

/**
 * This is the base interface of all processes (executors). When a process is started,
 * a new thread is created to call the RunOnce() function of the derived class.
 * Each process also have an incoming event queue and an outgoing event queue.
 */
type Process interface {
	NewProcess()

	// Start the process.
	Start()

	// Run process once return true if the thread should continue; false if the thread should exist.
	RunOnce() bool
}

/**
 * The base class for executors of source and operator.
 */
type ComponentExecutor interface {
	Component
	Process

	// Get the instance executors of this component executor.
	GetInstanceExecutors() []InstanceExecutor

	SetIncomings(queues []*EventQueue)

	SetOutgoing(queue *EventQueue)

	Start()
}

/**
 * Due to the need to achieve concurrency,
 * InstanceExecutor takes on some of ComponentExecutor's responsibilities in v0.1
 */
type InstanceExecutor interface {
	Process

	SetIncoming(in *EventQueue)

	SetOutgoing(out *EventQueue)
}

/**
 * This is the base class for all the event classes.
 * Users should extend this class to implement all their own event classes.
 */
type Event interface {
	// Get data stored in the event.
	GetData() any
}

/**
 * This is the class for intemediate event queues between processes.
 */
type EventQueue struct {
	Queue chan Event
}
