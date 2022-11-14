package fraud_detection_job

import (
	"fmt"
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/operator"
)

type AvgTicketAnalyzer struct {
	operator.Operator
}

// AvgTicketAnalyzer is a dummy analyzer. Allow all transactions.
func NewAvgTicketAnalyzer(name string, args ...any) *AvgTicketAnalyzer {
	v := new(AvgTicketAnalyzer)

	v.Name = name
	switch len(args) {
	case 0:
		v.Parallelism = 1
	case 1:
		v.Parallelism = args[0].(int)
	case 2:
		v.Parallelism = args[0].(int)
		v.Strategy = (args[1].(engine.GroupStrategy)) // default strategy is round-robin
	default:
		panic("too many arguments for NewAvgTicketAnalyzer")
	}
	return v
}

func (a *AvgTicketAnalyzer) Apply(e engine.Event, ev engine.EventCollector) error {
	t, ok := e.(*TransactionEvent)
	if !ok {
		panic("should be transactionEvent")
	}
	ev.Add(NewTransactionSorceEvent(t, 0.0))
	fmt.Printf("")
	return nil
}
