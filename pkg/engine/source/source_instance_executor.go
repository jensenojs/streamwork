package source

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/component"
)

func NewSourceExecutorInstance(postBase, id int, so engine.Source) *SourceInstanceExecutor {
	var soi = new(SourceInstanceExecutor)
	soi.InstanceId = id
	soi.source = so
	soi.EventCollector = component.NewEventCollector()
	soi.OutgoingMap = make(map[engine.Channel][]engine.EventQueue)

	soi.setupSocketReader(postBase + id)
	soi.SetRunOnce(soi.RunOnce)
	return soi
}

func (s *SourceInstanceExecutor) RunOnce() bool {
	// generate events
	buf, num := s.GetFromNet()
	fmt.Printf("%s:(%d) --> ", s.source.GetName(), s.InstanceId)
	s.source.GetEvents(buf, num, s.EventCollector)

	// emit out
	s.SendOutgoingEvent()

	// clean up event that executed
	s.EventCollector.Clear()

	return true
}

func (s *SourceInstanceExecutor) setupSocketReader(port int) {
	listener, err := net.Listen(ConnType, ConnHost+":"+strconv.FormatInt(int64(port), 10))
	if err != nil {
		panic(err)
	}
	s.Ln = listener
}

func (s *SourceInstanceExecutor) GetFromNet() ([]byte, int) {
	buf := make([]byte, 1024)
	if s.Conn == nil {
		Conn, err := s.Ln.Accept()
		if err != nil {
			panic(err)
		}
		s.Conn = Conn
	}
	num, err := s.Conn.Read(buf)
	if err != nil {
		// disconnecting from client, for now just exit
		os.Exit(0)
	}
	return buf, num
}
