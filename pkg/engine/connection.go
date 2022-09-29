package engine

/**
 * A util data class for connections between components.
 */
type Connection struct {
	from ComponentExecutor // is interface because could be SourceExecutor or OperatorExecutor
	to   *OperatorExecutor
}

func NewConnection(from ComponentExecutor, to *OperatorExecutor) *Connection {
	return &Connection{
		from: from,
		to:   to,
	}
}
