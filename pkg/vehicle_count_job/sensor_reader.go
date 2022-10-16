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
	connPort = "" // no use
)

type SensorReader struct {
	engine.SourceExecutor

	// listener net.Listener
	conn net.Conn
}

func NewSensorReader(name string, port int64) *SensorReader {
	var s = &SensorReader{}
	// s.InitNameAndStream(name)
	// s.setupSocketReader(port)
	return s
}

func (s *SensorReader) GetEvents(eventCollector *[]api.Event) {
	buf := make([]byte, 1024)
	num, err := s.conn.Read(buf)
	if err != nil {
		// disconnecting from client, for now just exit
		os.Exit(0)
	}
	vehicle := NewVehicleEvent(string(buf[:num - 1])) // the last character is '\n', so just ignore it
	*eventCollector = append(*eventCollector, vehicle)
	fmt.Printf("SensorReader --> %s\n", vehicle.GetData())
}

func (s *SensorReader) setupSocketReader(port int64) {
	fmt.Println("SensorReader begin to monitor")

	listener, err := net.Listen(connType, connHost+":"+strconv.FormatInt(port, 10))
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
