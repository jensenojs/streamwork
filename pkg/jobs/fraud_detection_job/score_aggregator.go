package fraud_detection

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/operator"
)

type ScoreAggregator struct {
	operator.Operator
	store *ScoreStorage
}

func NewScoreAggregator(name string, args ...any) *ScoreAggregator {
	var s = &ScoreAggregator{
		store: NewScoreStorage(),
	}

	s.Name = name
	switch len(args) {
	case 0:
		s.Parallelism = 1
	case 1:
		s.Parallelism = args[0].(int)
	case 2:
		s.Parallelism = args[0].(int)
		s.Strategy = (args[1].(engine.GroupStrategy)) // default strategy is round-robin
	default:
		panic("too many arguments for NewVehicleCounter")
	}
	return s
}

func (s *ScoreAggregator) Apply(score engine.Event, ev engine.EventCollector) error {
	e, ok := score.(*TransactionScoreEvent)
	if !ok {
		panic("should be a transactionScoreEvent")
	}
	old := s.store.get(e.Tran.Id, 0)
	s.store.set(e.Tran.Id, old+e.Score)
	return nil
}
