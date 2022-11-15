package fraud_detection

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/transport/strategy"
)

type TranIdFieldStrategy struct {
	strategy.FieldGrouping
}

func NewTranIdFieldStrategy() *TranIdFieldStrategy {
	var tfs = new(TranIdFieldStrategy)
	tfs.Map = make(map[string]int)
	tfs.CustomGetKey = tfs.GetKey
	return tfs
}

func (tr *TranIdFieldStrategy) GetKey(e engine.Event) string {
	t := e.(*TransactionScoreEvent)
	return t.Tran.Id
}
