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

// Component is a interface for all components, including Source and Operator.
type Component interface {
	GetName() string

	// Get the outgoing event stream of this component. The stream is used to connect the downstream components.
	GetOutgoingStream() Stream

	// Get the parallelism (number of instances) of this component.
	GetParallelism() int
}

// Operator is a interface for all user defined operators.
type Operator interface {
	Component

	//  Apply logic to the incoming event and generate results.The function is abstract and needs to be implemented by users.
	Apply(Event, EventCollector) error

	// If you don't want to set GroupStrategy, just return nil
	GetGroupingStrategy() GroupStrategy
}

// Source is a interface for all user defined sources.
type Source interface {
	Component

	// Accept events from external into the system.  The function is abstract and needs to be implemented by users.
	GetEvents([]byte, int, EventCollector)
}

// ComponentExecutor is a interface for executors of source and operator. Executors is not a component
// but each component need a executor, as component executors implement process interface, which defined how process work in streamwork
type ComponentExecutor interface {
	process.Process

	// Get the instance executors of this component executor.
	GetInstanceExecutors() []InstanceExecutor

	SetIncomings([]EventQueue)

	AddOutgoing(Channel, EventQueue)

	RegisterChannel(Channel)
}

// InstanceExecutor take charge of specific work
// ComponentExecutor and InstanceExecutor may be in a one-to-one or one-to-many relationship (parallel)
type InstanceExecutor interface {
	process.Process

	SetIncoming(EventQueue)

	AddOutgoing(Channel, EventQueue)

	RegisterChannel(Channel)
}

// Get target instance id from an event and component parallelism.
// Note that in this implementation, only one instance is selected.
// This can be easily extended if needed.
type GroupStrategy interface {
	// the event object to route to the component
	// the parallelism of the component
	GetInstance(event Event, parallelism int) int
}

// This is the base class for all the event classes.
// Users should extend this class to implement all their own event classes.
type Event interface {
	// Get data stored in the event.
	IsEvent()

	// This method used to support filed_grouping
	GetKey() string
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
