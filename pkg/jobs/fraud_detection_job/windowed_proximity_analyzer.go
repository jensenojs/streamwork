package fraud_detection

import (
	"fmt"
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/operator"
)

// WindowedProximityAnalyzer is a dummy analyzer. Allow all transactions.
type WindowedProximityAnalyzer struct {
	operator.Operator
}

func NewWindowedProximityAnalyzer(name string, args ...any) *WindowedProximityAnalyzer {
	w := new(WindowedProximityAnalyzer)

	w.Name = name
	switch len(args) {
	case 0:
		w.Parallelism = 1
	case 1:
		w.Parallelism = args[0].(int)
	case 2:
		w.Parallelism = args[0].(int)
		w.Strategy = (args[1].(engine.GroupStrategy)) // default strategy is round-robin
	default:
		panic("too many arguments for NewWindowedProximityAnalyzer")
	}
	return w
}

func (w *WindowedProximityAnalyzer) Apply(e engine.Event, ev engine.EventCollector) error {
	t, ok := e.(*TransactionEvent)
	if !ok {
		panic("should be transactionEvent")
	}
	ev.Add(NewTransactionSorceEvent(t, 0.0))
	fmt.Printf("0.0")
	return nil
}
