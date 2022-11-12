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
		queue_size:       64, // default queue size
		operatorMap:      make(map[engine.Operator]*operator.OperatorExecutor),
		operatorQueueMap: make(map[*operator.OperatorExecutor]*transport.EventQueue),
		job:              job,
	}
}

func (j *JobStarter) Start() {
	// set up executors for all the components.
	j.setupComponentExecutors()

	// all components are created now. Build the connections to connect the components together.
	j.setupConnections()

	// start all the processes.
	j.startProcesses()

	// let main.go running
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
func (j *JobStarter) traverseComponent(from engine.Component, fromExe engine.ComponentExecutor) {
	downstream := from.GetOutgoingStream().(*stream.Stream)
	for _, ch := range downstream.GetChannels() {
		var toExe *operator.OperatorExecutor
		// get the operators apply on upstream components specific by channel
		for to := range downstream.GetAppliedOperators(ch) {
			if _, ok := j.operatorMap[to]; !ok {
				toExe = operator.NewOperatorExecutor(to)
				j.operatorMap[to] = toExe
				j.executorList = append(j.executorList, toExe)
				j.traverseComponent(to, toExe) // set up executors for the downstream operators
			} else {
				toExe = j.operatorMap[to]
			}
			// setup executors for the downstream operators
			j.connectionList = append(j.connectionList, NewConnection(fromExe, toExe, ch))
		}
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
func (j *JobStarter) connectExecutors(connection *Connection) {

	connection.From.RegisterChannel(connection.Channel)
	if q, ok := j.operatorQueueMap[connection.To]; ok {
		// Existing operator. Connect to upstream only.
		connection.From.AddOutgoing(connection.Channel, q)
	} else {
		// New operator. Create a dispatcher and connect to upstream first.
		d := transport.NewEventDispatcher(connection.To)
		dq := transport.NewEventQueue(j.queue_size)
		j.operatorQueueMap[connection.To] = dq
		d.SetIncoming(dq)
		connection.From.AddOutgoing(connection.Channel, dq)

		// connect to downstream (to each instance)
		p := connection.To.GetParallelism()
		ds := make([]engine.EventQueue, p)
		for i := range ds {
			ds[i] = transport.NewEventQueue(j.queue_size)
		}
		connection.To.SetIncomings(ds)
		d.SetOutgoings(ds)

		j.dispatcherList = append(j.dispatcherList, d)
	}
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

// reverseExecutorList reverseExecutor, If it starts from upstream, it may cause the loss of some events.
func (j *JobStarter) reverseExecutorList() {
	reverseExecutorList := []engine.ComponentExecutor{}
	for i := range j.executorList {
		reverseExecutorList = append(reverseExecutorList, j.executorList[len(j.executorList)-1-i])
	}
	j.executorList = reverseExecutorList
}
