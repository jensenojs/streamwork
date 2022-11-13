package fraud_detection_job

import (
	"fmt"
	"github.com/google/uuid"
	"strconv"
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/source"
	"strings"
)

type TransationSource struct {
	source.Source
}

func NewTransactionSource(name string, args ...any) *TransationSource {
	s := new(TransationSource)
	s.Name = name

	switch len(args) {
	case 0:
		s.Parallelism = 1
	case 1:
		s.Parallelism = args[0].(int)
	case 2:
		s.Parallelism = args[0].(int)
		s.Clone = args[1].(bool)
	default:
		panic("too many arguments for NewSensorReader")
	}

	return s
}

func (tr *TransationSource) GetEvents(t string, e engine.EventCollector) {
	var amount float64
	var merchandiseId int
	var err error

	values := strings.Split(t, ",")
	if values[0] == t {
		fmt.Printf("Input needs to be in this format: {amount},{merchandiseId}. For example: 42.00,3\n")
		return // No transaction to emit.
	} else {
		if amount, err = strconv.ParseFloat(values[0], 64); err != nil {
			fmt.Printf("Input amount is illegal, please double check your input")
			return
		}
		if merchandiseId, err = strconv.Atoi(values[1]); err != nil {
			fmt.Printf("Input merchandiseId is illegal, please double check your input")
			return
		}
	}

	// For simplicity, assuming all transactions are from the same user. Transaction id and time are generated automatically.
	userAccount := 1
	tid := uuid.New().String()

	event := NewTransactionEvent(tid, amount, merchandiseId, userAccount)
	e.Add(event)
	fmt.Printf("amount: %f, merchandiseId: %d\n", amount, merchandiseId)
}
