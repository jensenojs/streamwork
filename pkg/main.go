package main

import (
	"fmt"
	"streamwork/pkg/engine/job"
	"streamwork/pkg/engine/stream"
	"streamwork/pkg/jobs/fraud_detection_job"
)

// fraud detection base test
func main() {

	fraudJob := job.NewJob("fraud detection base test")

	transactionOut, err := fraudJob.AddSource(fraud_detection.NewTransactionSource("transaction source"))
	if err != nil {
		panic(err)
	}

	// One stream can have multiple channels. Different operator can be hooked up
	// to different channels to receive different events. When no channel is selected,
	// the default channel will be used.
	// All Analyzer are dummy, just pass all transactions, this job just just use to show stream fan-in and fan-out
	evalResult1, err := transactionOut.ApplyOperator(fraud_detection.NewAvgTicketAnalyzer("avg ticket analyzer", 2, fraud_detection.NewCarFiledStrategy()))

	evalResult2, err := transactionOut.ApplyOperator(fraud_detection.NewWindowedProximityAnalyzer("windowed proximity analyzer", 2, fraud_detection.NewCarFiledStrategy()))

	evalResult3, err := transactionOut.ApplyOperator(fraud_detection.NewWindowedTransactionCountAnalyzer("windowed transaction count analyzer", 2, fraud_detection.NewCarFiledStrategy()))

	stream.Of(evalResult1, evalResult2, evalResult3).ApplyOperator(fraud_detection.NewScoreAggregator("score aggregator", 2, fraud_detection.NewTranIdFieldStrategy()))

	fmt.Println("This is a streaming job that detect suspicious transactions." + "\n" +
		"Input needs to be in this format: {amount},{merchandiseId}. For example: 42.00,3" + "\n" +
		"Merchandises N and N + 1 are 1 seconds walking distance away from each other.")

	starter := job.NewJobStarter(fraudJob)
	starter.Start()
}
