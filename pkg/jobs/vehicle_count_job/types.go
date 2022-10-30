package vehicle_count_job

import (
	"net"
	"streamwork/pkg/engine/operator"
	"streamwork/pkg/engine/source"
)

type carType = string

type VehicleEvent struct {
	Type carType
}

func (v *VehicleEvent) IsEvent() {}

// SensorReader is a monitor on the brigde, Track how many cars are passing by. specific to the type of the car
type SensorReader struct {
	source.SourceExecutor
	conn       net.Conn
	instanceId int
	portBase   int
	clone      bool
}

// VehicleCounter is a counter
type VehicleCounter struct {
	operator.OperatorExecutor
	counter    map[carType]int
	instanceId int
}
