package api

import "errors"

type void struct{}

var member void

/**
 * The Job class is used by users to set up their jobs and run.
 * Example:
 *   Job job = new Job("my_job");
 *   job.addSource(mySource)
 *      .applyOperator(myOperator);
 */
type Job struct {
	name      string
	sourceSet map[Source]void
}

func NewJob(name string) *Job {
	return &Job{
		name:      name,
		sourceSet: make(map[Source]void),
	}
}

func (j *Job) GetName() string {
	return j.name
}

/**
 * Get the list sources in this job. This function is used by JobRunner to traverse the graph.
 * @return The list of sources in this job
 */
func (j *Job) GetSources() map[Source]void {
	return j.sourceSet
}

/**
 * Add a source into the job. A stream is returned which will be used to connect to
 * other operators.
 * @param source The source object to be added into the job
 * @return A stream that can be used to connect to other operators
 */
func (j *Job) AddSource(source Source) (stream *Stream, err error) {
	if _, ok := j.sourceSet[source]; ok {
		err = errors.New("Source" + source.GetName() + " already exists")
		return
	}
	j.sourceSet[source] = member
	return source.GetOutgoingStream(), nil
}
