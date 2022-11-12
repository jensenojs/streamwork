package operator

import (
	"streamwork/pkg/engine"
)

// =================================================================
// implement for Operator

func (o *Operator) Apply(engine.Event, engine.EventCollector) error {
	panic("need to implement Apply")
}

func (v *Operator) GetGroupingStrategy() engine.GroupStrategy {
	return v.Strategy
}
