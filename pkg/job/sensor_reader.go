package job

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"streamwork/pkg/api"
	"streamwork/pkg/engine"
)

type SensorReader struct {
	engine.SourceExecutor

	// listener net.Listener
	conn net.Conn
}

func NewSensorReader(name string, port int64) *SensorReader {
	var s = &SensorReader{}
	s.InitNameAndStream(name)
	s.setupSocketReader(port)
	return s
}

func (s *SensorReader) GetEvents(eventCollector []api.Event) {
	buf := make([]byte, 1024)
	num, err := s.conn.Read(buf)
	if err != nil {
		if err == io.EOF {
			err = nil
		} else {
			panic(err)
		}
	}
	vehicle := NewVehicleEvent(string(buf[:num]))
	eventCollector = append(eventCollector, vehicle)
	fmt.Printf("SensorReader --> %s\n", vehicle.GetData())
}

func (s *SensorReader) setupSocketReader(port int64) {
	fmt.Println("SensorReader begin to monitor")
	var listener net.Listener
	var err error
	if listener, err = net.Listen("tcp", "localhost:"+strconv.FormatInt(port, 10)); err != nil {
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
