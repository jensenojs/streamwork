package vehicle_count_job

func NewVehicleEvent(Type string) *VehicleEvent {
	return &VehicleEvent{
		Type: Type,
	}
}

// implement for Event interface
func (v *VehicleEvent) GetKey() carType {
	return v.Type
}
