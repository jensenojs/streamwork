package vehicle_count_job

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"streamwork/pkg/api"
	"streamwork/pkg/engine"
)

const (
	connHost = "localhost"
	connType = "tcp"
	connPort = 9990
)

type SensorReader struct {
	engine.SourceExecutor

	// listener net.Listener
	conn net.Conn

	instanceId int
	portBase   int
}

func NewSensorReader(name string, args ...int) *SensorReader {
	var s = &SensorReader{}

	if len(args) == 0 {
		s.Init(name, 1)
		s.portBase = connPort
	} else if len(args) == 1 {
		s.Init(name, args[0])
		s.portBase = connPort
	} else if len(args) == 2 {
		s.Init(name, args[0])
		s.portBase = args[1]
	}

	if len(args) > 2 {
		panic("too many arguments for NewSensorReader")
	}

	return s
}

func (s *SensorReader) SetupInstance(instanceId int) {
	s.instanceId = instanceId
	s.setupSocketReader(s.portBase + s.instanceId)
}

func (s *SensorReader) GetEvents(eventCollector *[]api.Event) {
	buf := make([]byte, 1024)
	num, err := s.conn.Read(buf)
	if err != nil {
		// disconnecting from client, for now just exit
		os.Exit(0)
	}
	vehicle := NewVehicleEvent(string(buf[:num-1])) // the last character is '\n', so just ignore it
	*eventCollector = append(*eventCollector, vehicle)
	fmt.Printf("SensorReader(%d) --> %s\n", s.instanceId, vehicle.GetData())
}

func (s *SensorReader) setupSocketReader(port int) {
	fmt.Println("SensorReader begin to monitor")

	listener, err := net.Listen(connType, connHost+":"+strconv.FormatInt(int64(port), 10))
	if err != nil {
		panic(err)
	}

	// wait for client connection
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Accept client con=%v 客户端ip = %v\n", conn, conn.RemoteAddr().String())
	}
	s.conn = conn
}
