package job

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/operator"
)

// A util data class for connections between components.
type Connection struct {
	From    engine.ComponentExecutor // is interface because could be SourceExecutor or OperatorExecutor
	To      *operator.OperatorExecutor
	Channel engine.Channel
}

func NewConnection(from engine.ComponentExecutor, to *operator.OperatorExecutor, ch engine.Channel) *Connection {
	return &Connection{
		From:    from,
		To:      to,
		Channel: ch,
	}
}
