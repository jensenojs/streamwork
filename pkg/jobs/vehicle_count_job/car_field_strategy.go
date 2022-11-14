package vehicle_count_job

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/transport/strategy"
)

type CarFiledStrategy struct {
	strategy.FieldGrouping
}

func NewCarFiledStrategy() *CarFiledStrategy {
	var cfs = new(CarFiledStrategy)
	cfs.Map = make(map[string]int)
	cfs.CustomGetKey = cfs.GetKey
	return cfs
}

func (c *CarFiledStrategy) GetKey(e engine.Event) string {
	ct := e.(*VehicleEvent)
	return ct.Type
}
