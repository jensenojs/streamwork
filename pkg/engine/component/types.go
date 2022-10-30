package component

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/stream"
)

// ComponentExecutorImpl used to Inherited by OperatorExecutor and SourceExecutor to save the implementation of some methods.
type ComponentExecutorImpl struct {
	Name              string
	Stream            *stream.Stream // connect to next component
	Parallelism       int
	InstanceExecutors []engine.InstanceExecutor
}

type InstanceExecutorImpl struct {
	InstanceId     int
	FnWrapper      func()                                 // wrapper function for fn to run continuously
	Fn             func() bool                            // process function runOnce, need to specific implementation for user logic
	EventCollector engine.EventCollector                  // help manage event dispatch
	Incoming       engine.EventQueue                      // for upstream processes
	OutgoingMap    map[engine.Channel][]engine.EventQueue // for downstream processes
}

type EventCollector struct {
	defaultChannel   engine.Channel
	List             map[engine.Channel][]engine.Event
	RegisterChannels map[engine.Channel]engine.Void
}
