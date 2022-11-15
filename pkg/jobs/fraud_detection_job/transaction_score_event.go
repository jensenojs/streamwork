package fraud_detection

// TransactionScoreEvent is a simple transaction event used in the fraud detection job.
type TransactionScoreEvent struct {
	Tran  *TransactionEvent
	Score float64
}

func (t *TransactionScoreEvent) IsEvent() {}

func (t *TransactionScoreEvent) GetKey() string {
	return t.Tran.Id
}

func NewTransactionSorceEvent(t *TransactionEvent, s float64) *TransactionScoreEvent {
	return &TransactionScoreEvent{
		Tran:  t,
		Score: s,
	}
}
