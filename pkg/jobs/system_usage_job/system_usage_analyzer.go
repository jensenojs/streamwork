package system_usage

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/operator"
)

type SystemUsageAnalyzer struct {
	operator.Operator
	transactionCount      int
	fraudTransactionCount int
}

func NewSystemUsageAnalyzer(name string, args ...any) *SystemUsageAnalyzer {
	var s = new(SystemUsageAnalyzer)

	s.Name = name
	switch len(args) {
	case 0:
		s.Parallelism = 1
	case 1:
		s.Parallelism = args[0].(int)
	case 2:
		s.Parallelism = args[0].(int)
		s.Strategy = (args[1].(engine.GroupStrategy)) // default strategy is round-robin
	default:
		panic("too many arguments for NewVehicleCounter")
	}
	return s
}

func (s *SystemUsageAnalyzer) Apply(score engine.Event, ev engine.EventCollector) error {
	s.transactionCount++
	// TODO: uncomment the code below to count fraud transactions.
	// TransactionEvent e = ((TransactionEvent)event);
	// String id = ((TransactionEvent)event).transactionId;
	// Thread.sleep(20);
	// boolean fraud = fraudStore.getItem(id);
	// if (fraud) {
	//        fraudTransactionCount++;
	//}
	ev.Add(NewUsageEvent(s.transactionCount, s.fraudTransactionCount))
	return nil
}
