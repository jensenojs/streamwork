package fraud_detection

import (
	"strconv"
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/transport/strategy/groupstrategy"
)

// define another filed_grouping strategy here, using userAccount but not Id
type UserAccountFieldStrategy struct {
	groupstrategy.FieldGrouping
}

func NewUserAccountStrategy() *UserAccountFieldStrategy {
	var ufs = new(UserAccountFieldStrategy)
	ufs.Map = make(map[string]int)
	ufs.CustomGetKey = ufs.GetKey
	return ufs
}

func (c *UserAccountFieldStrategy) GetKey(e engine.Event) string {
	t := e.(*TransactionEvent)
	return strconv.Itoa(t.UserAccount)
}
