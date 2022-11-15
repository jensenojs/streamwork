package vehicle_count

type carType = string

type VehicleEvent struct {
	Type carType
}

func (v *VehicleEvent) IsEvent() {}

func NewVehicleEvent(Type string) *VehicleEvent {
	return &VehicleEvent{
		Type: Type,
	}
}
