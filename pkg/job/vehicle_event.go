package job

type carType = string

type VehicleEvent struct {
	Type carType
}

func NewEventQueue(Type string) *VehicleEvent {
	return &VehicleEvent{
		Type : Type,
	}
}

// implement for Event interface
func (v *VehicleEvent) GetData() any {
	return v.Type
}