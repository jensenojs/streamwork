package job

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/operator"
	"streamwork/pkg/engine/transport"
)

// The Job class is used by users to set up their jobs and run.
// Example:
//
//	Job job = new Job("my");
//	job.addSource(mySource)
//	   .applyOperator(myOperator);
type Job struct {
	name      string
	sourceSet map[engine.Source]void
}

type JobStarter struct {
	queue_size int  // default queue size is 64
	job        *Job // the job to start

	executorList     []engine.ComponentExecutor                     // list of executors
	connectionList   []*Connection                                  // connections between component executors
	dispatcherList   []*transport.EventDispatcher                   // dispatcher is to support dispatch strategy
	operatorMap      map[engine.Operator]*operator.OperatorExecutor // note the different from stream's OperatorMap
	operatorQueueMap map[*operator.OperatorExecutor]*transport.EventQueue
}
