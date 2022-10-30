package job

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/operator"
	"streamwork/pkg/engine/process"
	"streamwork/pkg/engine/source"
	"streamwork/pkg/engine/stream"
	"streamwork/pkg/engine/transport"
)

func NewJobStarter(job *Job) *JobStarter {
	return &JobStarter{
		queue_size: 64, // default queue size
		job:        job,
	}
}

func (j *JobStarter) Start() {
	// set up executors for all the components.
	j.setupComponentExecutors()

	// all components are created now. Build the connections to connect the components together.
	j.setupConnections()

	// start all the processes.
	j.startProcesses()

	// let main.go running forever
	for {

	}

}

// =================================================================
// setupComponentExecutors setup ComponentExecutors and helper functions
func (j *JobStarter) setupComponentExecutors() {
	// start from sources in the job and traverse components to create executors
	for s := range j.job.GetSources() {
		se := source.NewSourceExecutor(s)
		// for each source, traverse the operations connected to it
		j.executorList = append(j.executorList, se)
		j.traverseComponent(s, se) // traverse begin with upstream
	}
}

// traverseComponent traverse the components from source and initialize them
func (j *JobStarter) traverseComponent(from engine.Component, fromExecutor engine.ComponentExecutor) {
	downstream := from.GetOutgoingStream().(*stream.Stream)
	// get the operators apply on upstream components
	for t := range downstream.GetAppliedOperators() {
		te := operator.NewOperatorExecutor(t)
		j.executorList = append(j.executorList, te)
		j.connectionList = append(j.connectionList, transport.NewConnection(fromExecutor, te))
		// setup executors for the downstream operators
		j.traverseComponent(t, te)
	}
}

// =================================================================
// setupConnections use connectExecutors to connect component in this job
func (j *JobStarter) setupConnections() {
	for _, c := range j.connectionList {
		j.connectExecutors(c)
	}
}

// connectExecutors Each component executor could connect to multiple downstream operator executors.
// For each of the downstream operator executor, there is a stream manager.
// Each instance executor of the upstream executor connects to the all the stream managers
// of the downstream executors first. And each stream manager connects to all the instance
// executors of the downstream executor.
// Note that in this version, there is no shared "from" component and "to" component.
// The job looks like a single linked list.
func (j *JobStarter) connectExecutors(connection *transport.Connection) {
	d := NewEventDispatcher(connection.To)
	j.dispatcherList = append(j.dispatcherList, d)

	// connect to upstream
	upstream := transport.NewEventQueue(j.queue_size)
	connection.From.SetOutgoing(upstream)
	d.SetIncoming(upstream)

	// connect to downstream (to each instance)
	p := connection.To.GetParallelism()
	ds := make([]*engine.EventQueue, p)
	for i := range ds {
		ds[i] = transport.NewEventQueue(j.queue_size)
	}
	connection.To.SetIncomings(ds)
	d.SetOutgoings(ds)
}

// =================================================================
// startProcess start all the processes for this job and helper function
func (j *JobStarter) startProcesses() {
	j.reverseExecutorList()
	for _, e := range j.executorList {
		e.(process.Process).NewProcess()
		e.(process.Process).Start()
	}
	for _, d := range j.dispatcherList {
		d.NewProcess()
		d.Start()
	}
}

// reverseExecutorList reverseExecutor
func (j *JobStarter) reverseExecutorList() {
	reverseExecutorList := []engine.ComponentExecutor{}
	for i := range j.executorList {
		reverseExecutorList = append(reverseExecutorList, j.executorList[len(j.executorList)-1-i])
	}
	j.executorList = reverseExecutorList
}
