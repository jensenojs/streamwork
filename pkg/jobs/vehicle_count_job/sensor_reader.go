package vehicle_count_job

import (
	"fmt"
	"net"
	"os"

	"strconv"
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

func (s *SensorReader) SetupInstance(instanceId int) {
	s.instanceId = instanceId
	s.setupSocketReader(s.portBase + s.instanceId)
}

func (s *SensorReader) GetEvents(eventCollector engine.EventCollector) {
	buf := make([]byte, 1024)
	if s.conn == nil {
		conn, err := s.ln.Accept()
		if err != nil {
			panic(err)
		}
		s.conn = conn
	}
	num, err := s.conn.Read(buf)
	if err != nil {
		// disconnecting from client, for now just exit
		os.Exit(0)
	}
	vehicle := string(buf[:num-1])
	// This source emits events into two channels.
	eventCollector.Add(NewVehicleEvent(vehicle))
	if s.clone {
		eventCollector.Addto(NewVehicleEvent(vehicle+"-clone"), "clone")
	}
	fmt.Printf("%s:(%d) --> %s\n", s.GetName(), s.instanceId, vehicle)
}

func (s *SensorReader) setupSocketReader(port int) {
	listener, err := net.Listen(jobs.ConnType, jobs.ConnHost+":"+strconv.FormatInt(int64(port), 10))
	if err != nil {
		panic(err)
	}
	s.ln = listener
}
