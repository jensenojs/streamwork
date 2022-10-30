package job

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/operator"
)

// A util data class for connections between components.
type Connection struct {
	From engine.ComponentExecutor // is interface because could be SourceExecutor or OperatorExecutor
	To   *operator.OperatorExecutor
}

func NewConnection(from engine.ComponentExecutor, to *operator.OperatorExecutor) *Connection {
	return &Connection{
		From: from,
		To:   to,
	}
}
