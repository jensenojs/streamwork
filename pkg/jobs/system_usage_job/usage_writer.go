package system_usage

import (
	"fmt"
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/operator"
)

type UsageWriter struct {
	operator.Operator
}

func NewUsageWriter(name string, args ...any) *UsageWriter {
	var s = new(UsageWriter)

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

func (u *UsageWriter) Apply(e engine.Event, ev engine.EventCollector) error {
	fmt.Print("writeing...")
	return nil
}
