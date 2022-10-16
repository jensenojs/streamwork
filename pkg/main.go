package main

import (
	"fmt"
	"streamwork/pkg/api"
	"streamwork/pkg/engine"
	"streamwork/pkg/vehicle_count_job"
)

// vehicle count job
func main() {
	vehicleJob := api.NewJob("vehicle count")
	brigdeStream, err := vehicleJob.AddSource(vehicle_count_job.NewSensorReader("sensor-reader", 2))
	if err != nil {
		panic(err)
	}
	brigdeStream.ApplyOperator(vehicle_count_job.NewVehicleCounter("vehicle counter", 3))

	fmt.Println("This is a streaming job that counts vehicles in real time. " +
		"Please enter vehicle types like 'car' and 'truck' in the input terminal " +
		"and look at the output")

	starter := engine.NewJobStarter(vehicleJob)
	starter.Start()
}
