package vehicle_count_job

import (
	"fmt"
	"sort"
	"streamwork/pkg/api"
	"streamwork/pkg/engine"
)

type VehicleCounter struct {
	engine.OperatorExecutor
	counter    map[carType]int
	instanceId int
}

func NewVehicleCounter(name string, args ...any) *VehicleCounter {
	var v = &VehicleCounter{
		counter: make(map[carType]int),
		instanceId: 0,
	}
	
	switch len(args) {
	case 0:
		v.Init(name, 1)
	case 1:
		v.Init(name, args[0].(int))
	case 2:
		v.Init(name, args[0].(int))
		v.SetGroupingStrategy(args[1].(api.GroupStrategy)) // in fact, default strategy is round-robin
	default:
		panic("too many arguments for NewVehicleCounter")
	}
	return v
}

func (v *VehicleCounter) SetupInstance(instanceId int) {
	v.instanceId = instanceId
}

func (v *VehicleCounter) Apply(vehicleEvent api.Event, eventCollector *[]api.Event) error {
	vehicle := vehicleEvent.(*VehicleEvent).GetData().(carType)
	v.counter[vehicle] = v.counter[vehicle] + 1

	fmt.Printf("VehicleCounter(%d) --> ", v.instanceId)
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
