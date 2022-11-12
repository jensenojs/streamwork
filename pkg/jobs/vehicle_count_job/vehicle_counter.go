package vehicle_count_job

import (
	"fmt"
	"sort"
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/transport/strategy"
)

func NewVehicleCounter(name string, args ...any) *VehicleCounter {
	var v = &VehicleCounter{
		counter: make(map[carType]int),
	}

	switch len(args) {
	case 0:
		v.Init(name, 1)
	case 1:
		v.Init(name, args[0].(int))
	case 2:
		v.Init(name, args[0].(int))
		v.SetGroupingStrategy(args[1].(strategy.GroupStrategy)) // in fact, default strategy is round-robin
	default:
		panic("too many arguments for NewVehicleCounter")
	}
	return v
}

// =================================================================
// implement for Operator

func (v *VehicleCounter) Apply(e engine.Event, _ engine.EventCollector) error {
	typ := e.(*VehicleEvent).GetType()
	v.counter[typ] = v.counter[typ] + 1
	v.printCountMap()
	return nil
}

func (v *VehicleCounter) printCountMap() {
	keys := make([]carType, 0)
	for k := range v.counter {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("  "+"%s : "+"%d\n", k, v.counter[k])
	}
}
