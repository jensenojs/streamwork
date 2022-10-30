package job

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/transport"
)

// The Job class is used by users to set up their jobs and run.
// Example:
//
//	Job job = new Job("my_job");
//	job.addSource(mySource)
//	   .applyOperator(myOperator);
type Job struct {
	name      string
	sourceSet map[engine.Source]void
}

type JobStarter struct {
	queue_size     int
	job            *Job                       // the job to start
	executorList   []engine.ComponentExecutor // list of executors
	connectionList []*transport.Connection    // connections between component executors
	dispatcherList []*EventDispatcher
}
