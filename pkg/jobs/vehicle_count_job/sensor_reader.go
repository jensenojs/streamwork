package vehicle_count_job

import (
	"fmt"
	"streamwork/pkg/engine"
	"streamwork/pkg/jobs"
)

func NewSensorReader(name string, args ...any) *SensorReader {
	var s = &SensorReader{}

	switch len(args) {
	case 0:
		s.Init(name, 1)
		s.portBase = jobs.ConnPort
	case 1:
		s.Init(name, args[0].(int))
		s.portBase = jobs.ConnPort
	case 2:
		s.Init(name, args[0].(int))
		s.portBase = args[1].(int)
	case 3:
		s.Init(name, args[0].(int))
		s.portBase = args[1].(int)
		s.clone = args[2].(bool)
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
	if s.clone {
		e.Addto(NewVehicleEvent(v+"-clone"), "clone")
	}
	fmt.Printf("%s\n", v)
}
