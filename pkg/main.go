package main

import (
	"fmt"
	"streamwork/pkg/api"
	"streamwork/pkg/engine"
	"streamwork/pkg/job"
)

// vehicle count job
func main() {
	vehicleJob := api.NewJob("vehicle count")
	brigdeStream, err := vehicleJob.AddSource(job.NewSensorReader("sensor-reader", 9990))
	if err != nil {
		panic(err)
	}
	brigdeStream.ApplyOperator(job.NewVehicleCounter("vehicle counter"))

	fmt.Println("This is a streaming job that counts vehicles in real time. " +
		"Please enter vehicle types like 'car' and 'truck' in the input terminal " +
		"and look at the output")

	starter := engine.NewJobStarter(vehicleJob)
	starter.Start()
}
