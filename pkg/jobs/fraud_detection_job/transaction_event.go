package fraud_detection_job

import (
	"fmt"
	"time"
)

type TransactionEvent struct {
	Id            string
	Amount        float64
	Time          time.Time
	MerchandiseId int
	UserAccount   int
}

func (t *TransactionEvent) IsEvent() {}

func (t *TransactionEvent) GetKey() string {
	return t.Id
}

func NewTransactionEvent(id string, amount float64, mid, userAcc int) *TransactionEvent {
	return &TransactionEvent{
		Id:            id,
		Amount:        amount,
		Time:          time.Now(),
		MerchandiseId: mid,
		UserAccount:   userAcc,
	}
}

func (t *TransactionEvent) String() string {
	return fmt.Sprintf("Transaction: %s amount: %f transactionTime %s \n	merchandise: %d , user: %d",
		t.Id, t.Amount, t.Time.String(), t.MerchandiseId, t.UserAccount)
}
