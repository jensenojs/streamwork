package vehicle_count_job

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
