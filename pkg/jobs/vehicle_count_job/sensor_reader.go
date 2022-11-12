package vehicle_count_job

import (
	"fmt"
	"streamwork/pkg/engine"
)

func NewSensorReader(name string, args ...any) *SensorReader {
	s := new(SensorReader)
	s.Name = name

	switch len(args) {
	case 0:
		s.Parallelism = 1
	case 1:
		s.Parallelism = args[0].(int)
	case 2:
		s.Parallelism = args[0].(int)
		s.Clone = args[1].(bool)
	default:
		panic("too many arguments for NewSensorReader")
	}

	return s
}

// =================================================================
// implement for Source

func (s *SensorReader) GetEvents(buf []byte, num int, e engine.EventCollector) {
	// This source emits events into two channels.
	v := string(buf[:num-1])
	e.Add(NewVehicleEvent(v))
	if s.Clone {
		e.Addto(NewVehicleEvent(v+"-clone"), "clone")
	}
	fmt.Printf("%s\n", v)
}
