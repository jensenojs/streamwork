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

// =================================================================
// implement for WindowOperator

func (o *WindowOperator) Apply(engine.EventWindow, engine.EventCollector) error {
	panic("need to implement Apply")
}

func (v *WindowOperator) GetWindowingStrategy() engine.WindowStrategy {
	return v.Strategy
}
