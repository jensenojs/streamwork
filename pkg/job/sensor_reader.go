package job

import (
	"bytes"
	"streamwork/pkg/api"
	"streamwork/pkg/engine"
)

type SensorReader struct {
	engine.SourceExecutor

	reader bytes.Buffer
}

func NewSensorReader(name string, port int) *SensorReader {
	return &SensorReader{
		name : name,
	}
	return nil
}

func (s *SensorReader) GetEvents(eventCollector []api.Event) error {
	return nil
}

func (s *SensorReader) setupSocketReader(port int) {

}