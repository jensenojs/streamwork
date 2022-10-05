package vehicle_count_job

type carType = string

type VehicleEvent struct {
	Type carType
}

func NewVehicleEvent(Type string) *VehicleEvent {
	return &VehicleEvent{
		Type : Type,
	}
}

// implement for Event interface
func (v *VehicleEvent) GetData() any {
	return v.Type
}