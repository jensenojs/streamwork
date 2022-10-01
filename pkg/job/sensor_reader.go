package job

import (
	"bytes"
	"streamwork/pkg/api"
	"streamwork/pkg/engine"
)

type SensorReader struct {
	engine.SourceExecutor

	reader bytes.Reader
}

func NewSensorReader(name string, port int) *SensorReader {
	var s = &SensorReader{}
	s.Init(name) // where is port?
	s.setupSocketReader(port)
	return s
}

func (s *SensorReader) GetEvents(eventCollector []api.Event) error {
	// vehicle := 
	return nil
}

func (s *SensorReader) setupSocketReader(port int) {

}