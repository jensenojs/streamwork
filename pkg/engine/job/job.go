package job

import (
	"errors"
	"streamwork/pkg/engine"
)

type void struct{}

var member void

func NewJob(name string) *Job {
	return &Job{
		name:      name,
		sourceSet: make(map[engine.Source]void),
	}
}

func (j *Job) GetName() string {
	return j.name
}

// GetSources Get the list sources in this job. This function is used by job runner to traverse the graph.
func (j *Job) GetSources() map[engine.Source]void {
	return j.sourceSet
}

// AddSource Add a source into the job. A stream is returned which will be used to connect to other operators.
func (j *Job) AddSource(source engine.Source) (any, error) {
	if _, ok := j.sourceSet[source]; ok {
		return nil, errors.New("Source" + source.GetName() + " already exists")
	}
	j.sourceSet[source] = member
	return source.GetOutgoingStream(), nil
}
