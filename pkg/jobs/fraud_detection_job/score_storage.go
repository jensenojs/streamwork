package fraud_detection_job

import "fmt"

type ScoreStorage struct {
	transactionScores map[string]float64
}

func NewScoreStorage() *ScoreStorage {
	return &ScoreStorage{
		transactionScores: make(map[string]float64),
	}
}

func (s *ScoreStorage) get(transaction string, defaultValue float64) float64 {
	if v, ok := s.transactionScores[transaction]; ok {
		return v
	} else {
		return defaultValue
	}
}

func (s *ScoreStorage) set(transaction string, Value float64) {
	fmt.Printf("Transaction score change : "+"%s"+" ==> "+"%f\n", transaction, Value)
	s.transactionScores[transaction] = Value
}
