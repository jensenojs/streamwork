package engine

import "streamwork/pkg/api"

type JobStarter struct {
	queue_size     int
	job            api.Job             // the job to start
	executorList   []ComponentExecutor // list of executors
	connectionList []*Connection       // connections between component executors
}

func NewJobStarter(job api.Job) *JobStarter {
	return &JobStarter{
		queue_size: 64, // default queue size
		job:        job,
	}
}

func (j *JobStarter) Start() error {
	// set up executors for all the components.
	j.setupComponentExecutors()

	// all components are created now. Build the connections to connect the components together.
	j.setupConnections()

	// start all the processes.
	j.startProcesses()

	// start web server
	//
	return nil
}

func (j *JobStarter) setupComponentExecutors() {
	// start from sources in the job and traverse components to create executors
	for s := range j.job.GetSources() {
		e := NewSourceExecutor(s)
		// for each source, traverse the operations connected to it
		j.executorList = append(j.executorList, e)
		j.traverseComponent(s, e)
	}
}

func (j *JobStarter) traverseComponent(component api.Component, excecutor ComponentExecutor) {
	s := component.GetOutgoingStream()

	for o := range s.GetAppliedOperators() {
		oe := NewOperatorExecutor(o)
		j.executorList = append(j.executorList, oe)
		j.connectionList = append(j.connectionList, NewConnection(excecutor, *oe))
		// setup executors for the downstream operators
		j.traverseComponent(o, oe)
	}
}

func (j *JobStarter) setupConnections() {
	for _, c := range j.connectionList {
		j.connectExecutors(*c)
	}
}

// It is a newly connected operator executor. Note that in this version, there is no
// shared "from" component and "to" component. The job looks like a single linked list.
func (j *JobStarter) connectExecutors(connection Connection) {
	intermediateQueue := NewEventQueue(j.queue_size)
	connection.from.SetOutgoingQueue(*intermediateQueue)
	connection.to.SetIncomingQueue(*intermediateQueue)
}

// start all the processes for this job
func (j *JobStarter) startProcesses() {
	j.reverseExecutorList()
	for _, e := range j.executorList {
		e.(*componentExecutor).Start()
	}
}

func (j *JobStarter) reverseExecutorList() {
	reverseExecutorList := []ComponentExecutor{}
	for i := range j.executorList {
		reverseExecutorList = append(reverseExecutorList, j.executorList[len(j.executorList)-1-i])
	}
	j.executorList = reverseExecutorList
}