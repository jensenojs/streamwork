package job

import (
	"fmt"
	"sort"
	"streamwork/pkg/api"
	"streamwork/pkg/engine"
)

type VehicleCounter struct {
	engine.OperatorExecutor
	counter map[carType]int
}

func NewVehicleCounter(name string) *VehicleCounter {
	var v = &VehicleCounter{}
	v.Init(name)
	return v
}

func (v *VehicleCounter) Apply(vehicleEvent api.Event, eventCollector []api.Event) error {
	vehicle := vehicleEvent.(*VehicleEvent).GetData().(carType)
	v.counter[vehicle] = v.counter[vehicle] + 1

	fmt.Println("VehicleCounter --> ")
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
		fmt.Printf("  " + "%s :" +  "%d", k, v.counter[k])
	}
}