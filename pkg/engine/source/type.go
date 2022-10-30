package source

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/component"
)

// The executor for source components. When the executor is started,
// a new thread is created to call the getEvents() function of
// the source component repeatedly.
//
// Used to inherited by specific operator
type SourceExecutor struct {
	component.ComponentExecutorImpl
	so engine.Source // specific source, used to execute GetEvent
}

type SourceInstanceExecutor struct {
	component.InstanceExecutorImpl
	source     engine.Source
}
