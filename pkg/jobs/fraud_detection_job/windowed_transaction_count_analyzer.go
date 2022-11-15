package fraud_detection

import (
	"fmt"
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/operator"
)

// WindowedTransactionCountAnalyzer is a dummy analyzer. Allow all transactions.
type WindowedTransactionCountAnalyzer struct {
	operator.Operator
}

func NewWindowedTransactionCountAnalyzer(name string, args ...any) *WindowedTransactionCountAnalyzer {
	w := new(WindowedTransactionCountAnalyzer)

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
		panic("too many arguments for NewWindowedTransactionCountAnalyzer")
	}
	return w
}

func (w *WindowedTransactionCountAnalyzer) Apply(e engine.Event, ev engine.EventCollector) error {
	t, ok := e.(*TransactionEvent)
	if !ok {
		panic("should be transactionEvent")
	}
	ev.Add(NewTransactionSorceEvent(t, 0.0))
	fmt.Printf("0.0")
	return nil
}
