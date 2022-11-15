package system_usage

import (
	"fmt"
	"streamwork/pkg/engine/job"
	"testing"
)

func TestBase(t *testing.T) {
	sysjob := job.NewJob("system usage job")

	transactionout, err := sysjob.AddSource(NewTransactionSource("transaction source"))
	if err != nil {
		panic(err)
	}

	evalResult, err := transactionout.ApplyOperator(NewSystemUsageAnalyzer("system usage analyzer", 2, NewTranIdFieldStrategy()))
	if err != nil {
		panic(err)
	}
	evalResult.ApplyOperator(NewUsageWriter("usage writer", 2))
	fmt.Println("This is a streaming job that detect suspicious transactions." + "\n" +
		"Input needs to be in this format: {amount},{merchandiseId}. For example: 42.00,3." + "\n" +
		"Merchandises N and N + 1 are 1 seconds walking distance away from each other.")

	starter := job.NewJobStarter(sysjob)
	starter.Start()
}
