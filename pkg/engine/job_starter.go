package engine

import "streamwork/pkg/api"

type JobStarter struct {
	queue_size     int
	job            *api.Job            // the job to start
	executorList   []ComponentExecutor // list of executors
	connectionList []*Connection       // connections between component executors
	dispatcherList []*EventDispatcher
}

func NewJobStarter(job *api.Job) *JobStarter {
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
// setup ComponentExecutors and helper functions
func (j *JobStarter) setupComponentExecutors() {
	// start from sources in the job and traverse components to create executors
	for source := range j.job.GetSources() {
		sourceExecutor := newSourceExecutor(source)
		// for each source, traverse the operations connected to it
		j.executorList = append(j.executorList, sourceExecutor)
		j.traverseComponent(source, sourceExecutor) // traverse begin with upstream
	}
}

func (j *JobStarter) traverseComponent(from api.Component, fromExecutor ComponentExecutor) {
	downstream := from.GetOutgoingStream()
	// get the operators apply on upstream components
	for to := range downstream.GetAppliedOperators() {
		toExecutor := newOperatorExecutor(to)
		j.executorList = append(j.executorList, toExecutor)
		j.connectionList = append(j.connectionList, NewConnection(fromExecutor, toExecutor))
		// setup executors for the downstream operators
		j.traverseComponent(to, toExecutor)
	}
}

// =================================================================
// setupConnections and helper function
func (j *JobStarter) setupConnections() {
	for _, c := range j.connectionList {
		j.connectExecutors(c)
	}
}

/**
 * Each component executor could connect to multiple downstream operator executors.
 * For each of the downstream operator executor, there is a stream manager.
 * Each instance executor of the upstream executor connects to the all the stream managers
 * of the downstream executors first. And each stream manager connects to all the instance
 * executors of the downstream executor.
 * Note that in this version, there is no shared "from" component and "to" component.
 * The job looks like a single linked list.
 */
func (j *JobStarter) connectExecutors(connection *Connection) {
	d := NewEventDispatcher(connection.to)
	j.dispatcherList = append(j.dispatcherList, d)

	// connect to upstream
	upstream := NewEventQueue(j.queue_size)
	connection.from.SetOutgoingQueue(upstream)
	d.SetIncomingQueue(upstream)

	// connect to downstream (to each instance)
	p := connection.to.GetParallelism()
	ds := make([]*EventQueue, p)
	for i := range ds {
		ds[i] = NewEventQueue(j.queue_size)
	}
	connection.to.SetIncomingQueues(ds)
	d.SetOutgoingQueues(ds)
}

// =================================================================
// start all the processes for this job and helper function
func (j *JobStarter) startProcesses() {
	j.reverseExecutorList()
	for _, e := range j.executorList {
		e.(Process).newProcess()
		e.(Process).Start()
	}
	for _, d := range j.dispatcherList {
		d.newProcess()
		d.Start()
	}
}

func (j *JobStarter) reverseExecutorList() {
	reverseExecutorList := []ComponentExecutor{}
	for i := range j.executorList {
		reverseExecutorList = append(reverseExecutorList, j.executorList[len(j.executorList)-1-i])
	}
	j.executorList = reverseExecutorList
}
