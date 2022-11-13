package fraud_detection_job

import (
	"fmt"
	"streamwork/pkg/engine/job"
	"streamwork/pkg/engine/stream"
	"testing"
)

func TestBase(t *testing.T) {
	fraudJob := job.NewJob("fraud detection base test")

	transactionOut, err := fraudJob.AddSource(NewTransactionSource("transaction source"))
	if err != nil {
		panic(err)
	}

	// One stream can have multiple channels. Different operator can be hooked up
	// to different channels to receive different events. When no channel is selected,
	// the default channel will be used.
	evalResult1, err := transactionOut.ApplyOperator(NewAvgTicketAnalyzer("avg ticket analyzer", 2))

	evalResult2, err := transactionOut.ApplyOperator(NewWindowedProximityAnalyzer("windowed proximity analyzer", 2))

	evalResult3, err := transactionOut.ApplyOperator(NewWindowedTransactionCountAnalyzer("windowed transaction count analyzer", 2))

	stream.Of(evalResult1, evalResult2, evalResult3).ApplyOperator(NewScoreAggregator("score aggregator", 2))

	fmt.Println("This is a streaming job that detect suspicious transactions." + "\n" +
		"Input needs to be in this format: {amount},{merchandiseId}. For example: 42.00,3" + "\n" +
		"Merchandises N and N + 1 are 1 seconds walking distance away from each other.")

	starter := job.NewJobStarter(fraudJob)
	starter.Start()
}
