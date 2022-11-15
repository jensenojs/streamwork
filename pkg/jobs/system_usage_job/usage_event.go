package system_usage

import "fmt"

type UsageEvent struct {
	transactionCount      int
	fraudTransactionCount int
}

func NewUsageEvent(t, f int) *UsageEvent {
	return &UsageEvent{
		transactionCount:      t,
		fraudTransactionCount: f,
	}
}

func (e *UsageEvent) IsEvent() {}

func (e *UsageEvent) String() string {
	return fmt.Sprintf("transaction count %d, fraud transaction count %d", e.transactionCount, e.fraudTransactionCount)
}
