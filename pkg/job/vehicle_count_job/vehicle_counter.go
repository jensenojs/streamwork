package vehicle_count_job

import (
	"fmt"
	"sort"
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/operator"
	"streamwork/pkg/engine/transport/strategy"
)

type VehicleCounter struct {
	operator.OperatorExecutor
	counter    map[carType]int
	instanceId int
}

func NewVehicleCounter(name string, args ...any) *VehicleCounter {
	var v = &VehicleCounter{
		counter:    make(map[carType]int),
		instanceId: 0,
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
func (v *VehicleCounter) SetupInstance(instanceId int) {
	v.instanceId = instanceId
}

func (v *VehicleCounter) Apply(vehicleEvent engine.Event, eventCollector *[]engine.Event) error {
	vehicle := vehicleEvent.(*VehicleEvent).GetData().(carType)
	v.counter[vehicle] = v.counter[vehicle] + 1

	fmt.Printf("VehicleCounter(%d) --> \n", v.instanceId)
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
