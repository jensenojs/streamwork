package vehicle_count

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/transport/strategy/groupstrategy"
)

type CarFiledStrategy struct {
	groupstrategy.FieldGrouping
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
