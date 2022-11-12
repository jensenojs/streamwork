package engine

import "streamwork/pkg/engine/process"

type Void struct{}

var Member Void

type Channel = string

const DEFAULT_CHANNEL = "default"

type Stream interface {
	ApplyOperator(Operator) (Stream, error)
	SelectChannel(Channel) Stream // need more docs here
}

/**
 * The base class for all components, including Source and Operator.
 */
type Component interface {
	GetName() string

	// Get the outgoing event stream of this component. The stream is used to connect the downstream components.
	GetOutgoingStream() Stream

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
	Apply(Event, EventCollector) error
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
	GetEvents([]byte, int, EventCollector)
}

// ComponentExecutor is a interface for executors of source and operator.
type ComponentExecutor interface {
	Component
	process.Process

	// Get the instance executors of this component executor.
	GetInstanceExecutors() []InstanceExecutor

	SetIncomings([]EventQueue)

	AddOutgoing(Channel, EventQueue)

	RegisterChannel(Channel)
}

/**
 * Due to the need to achieve concurrency,
 * InstanceExecutor takes on some of ComponentExecutor's responsibilities in v0.1
 */
type InstanceExecutor interface {
	process.Process

	SetIncoming(EventQueue)

	AddOutgoing(Channel, EventQueue)

	RegisterChannel(Channel)
}

/**
 * This is the base class for all the event classes.
 * Users should extend this class to implement all their own event classes.
 */
type Event interface {
	// Get data stored in the event.
	IsEvent()
}

// EventQueue is a interface for intemediate event queues between processes.
type EventQueue interface {
	Take() Event
	Send(Event)
}

// EventCollector is a field for component to help manage event dispatch
type EventCollector interface {
	GetRegisteredChannels() []Channel
	SetRegisterChannel(Channel)

	GetEventList(Channel) []Event

	Add(Event)
	Addto(Event, Channel)

	Clear()
}
