package system_usage

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/transport/strategy/groupstrategy"
)

type TranIdFieldStrategy struct {
	groupstrategy.FieldGrouping
}

func NewTranIdFieldStrategy() *TranIdFieldStrategy {
	var tfs = new(TranIdFieldStrategy)
	tfs.Map = make(map[string]int)
	tfs.CustomGetKey = tfs.GetKey
	return tfs
}

func (tr *TranIdFieldStrategy) GetKey(e engine.Event) string {
	t := e.(*TransactionEvent)
	return t.Id
}
