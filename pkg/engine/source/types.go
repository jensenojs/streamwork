package source

import (
	"net"
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/component"
)

const (
	ConnHost  = "localhost"
	ConnType  = "tcp"
	ConnPort  = 9990 //default port number
	ConnPort2 = 9990 + 10
)

type Source struct {
	component.Component
	Clone bool
}

// The executor for source components. When the executor is started,
// a new thread is created to call the getEvents() function of
// the source component repeatedly.
//
// Used to inherited by specific source
type SourceExecutor struct {
	component.ComponentExecutorImpl
	PortBase int
}

type SourceInstanceExecutor struct {
	component.InstanceExecutorImpl
	source engine.Source
	Ln     net.Listener
	Conn   net.Conn
	Clone  bool
}
