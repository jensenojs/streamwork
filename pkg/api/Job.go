package api

import "errors"


type Job struct {
	name string
	sourceSet map[Source]bool
}

func NewJob() *Job {
	return &Job{
		sourceSet: make(map[Source]bool),
	}
}

func (j *Job) GetName() string {
	return j.name
}

func (j *Job) GetSources() map[Source]bool {
	return j.sourceSet
}

func (j *Job) AddSource(source Source) (stream *Stream, err error) {
	if _, ok := j.sourceSet[source]; ok {
		err = errors.New("Source" + source.GetName() + " already exists")	
		return
	}
	j.sourceSet[source] = true
	return source.GetOutgoingStream(), nil
} 