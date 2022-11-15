// run as debug in this test so printf can be shown in debug console
package vehicle_count

import (
	"fmt"
	"streamwork/pkg/engine/job"
	"streamwork/pkg/engine/stream"
	"testing"
)

func TestBase(t *testing.T) {
	vehicleJob := job.NewJob("vehicle count base test")
	brigdeStream, err := vehicleJob.AddSource(NewSensorReader("sensor-reader"))
	if err != nil {
		panic(err)
	}
	brigdeStream.ApplyOperator(NewVehicleCounter("vehicle counter"))

	fmt.Println("This is a streaming job that counts vehicles in real time. " + "\n" +
		"Please enter vehicle types like 'car' and 'truck' in the input terminal " + "\n" +
		"and look at the output")

	starter := job.NewJobStarter(vehicleJob)
	starter.Start()
}

func TestParallelForSensor(t *testing.T) {
	vehicleJob := job.NewJob("vehicle count sensor parallel test ")
	brigdeStream, err := vehicleJob.AddSource(NewSensorReader("sensor-reader", 2))
	if err != nil {
		panic(err)
	}
	brigdeStream.ApplyOperator(NewVehicleCounter("vehicle counter"))

	fmt.Printf("This test is to test the parallelism of Sensor" + "\n" +
		"please open two terminals and connect with the corresponding port number respectively. " + "\n" +
		"the default port number is 9990 and 9991(e.g. nc localhost 9990; nc localhost 9991)\n")

	starter := job.NewJobStarter(vehicleJob)
	starter.Start()
}

func TestParallelForCounter(t *testing.T) {
	vehicleJob := job.NewJob("vehicle count counter parallel test")
	brigdeStream, err := vehicleJob.AddSource(NewSensorReader("sensor-reader"))
	if err != nil {
		panic(err)
	}
	brigdeStream.ApplyOperator(NewVehicleCounter("vehicle counter", 2))

	fmt.Printf("This test is to test the parallelism of Counter" + "\n" +
		"Please enter vehicle types like 'car' and 'truck' in the input terminal " + "\n" +
		"and look at the output" + "\n" +
		"the default port number is 9990(e.g. nc localhost 9990)\n")

	starter := job.NewJobStarter(vehicleJob)
	starter.Start()
}

func TestParallelBoth(t *testing.T) {
	vehicleJob := job.NewJob("vehicle count parallel test")
	brigdeStream, err := vehicleJob.AddSource(NewSensorReader("sensor-reader", 2))
	if err != nil {
		panic(err)
	}
	brigdeStream.ApplyOperator(NewVehicleCounter("vehicle counter", 3))

	fmt.Printf("This test is to test the parallelism of Both" + "\n" +
		"please open two terminals and connect with the corresponding port number respectively. " + "\n" +
		"the default port number is 9990(e.g. nc localhost 9990)\n")

	starter := job.NewJobStarter(vehicleJob)
	starter.Start()
}

// Belong tests use StreamChannel and Streams

func TestFork(t *testing.T) {
	vehicleJob := job.NewJob("vehicle count(fork)")
	brigdeStream, err := vehicleJob.AddSource(NewSensorReader("sensor-reader", 1, false))
	if err != nil {
		panic(err)
	}
	brigdeStream.ApplyOperator(NewVehicleCounter("vehicle counter(shuffle grouping)", 2))
	brigdeStream.ApplyOperator(NewVehicleCounter("vehicle counter(fields grouping)", 2, NewCarFiledStrategy()))

	fmt.Println("This is a streaming job that has two counting operators linked to " + "\n" +
		"the same input stream. One operator is configured with default " + "\n" +
		"grouping strategy (shuffle) and the other is configured with fields " + "\n" +
		"grouping strategy. Please enter vehicle types like 'car' and 'truck' in the " + "\n" +
		"input terminal and look at the output")

	starter := job.NewJobStarter(vehicleJob)
	starter.Start()
}

func TestSplit(t *testing.T) {
	vehicleJob := job.NewJob("vehicle count(split)")
	brigdeStream, err := vehicleJob.AddSource(NewSensorReader("sensor-reader", 1, true))
	if err != nil {
		panic(err)
	}
	brigdeStream.ApplyOperator(NewVehicleCounter("vehicle counter(shuffle grouping)", 2))
	brigdeStream.SelectChannel("clone").ApplyOperator(NewVehicleCounter("vehicle counter(fields grouping)", 2, NewCarFiledStrategy()))

	fmt.Println("This is a streaming job that has two counting operators linked to " + "\n" +
		"the same input stream. One operator is hooked up to the default channel of " + "\n" +
		"the stream and configured with default grouping strategy (shuffle). " + "\n" +
		"The other one is hooked up to the clone channel and configured with " + "\n" +
		"fields grouping strategy. Please enter vehicle types like 'car' and " + "\n" +
		"'truck' in the input terminal and look at the output")

	starter := job.NewJobStarter(vehicleJob)
	starter.Start()
}

func TestMerge(t *testing.T) {
	vehicleJob := job.NewJob("vehicle count(merge)")
	brigdeStream1, err := vehicleJob.AddSource(NewSensorReader("sensor-reader1", 1, false))
	if err != nil {
		panic(err)
	}
	brigdeStream2, err := vehicleJob.AddSource(NewSensorReader("sensor-reader2", 1, true))
	if err != nil {
		panic(err)
	}

	// This is the convenient way to apply an operator on multiple streams using
	// Streams.merge(...) function. The code works the same way as the following:
	//   Operator operator = new TollBooth("booth", 2, new FieldsGrouping());
	//   bridgeStream1.selectChannel("clone").applyOperator(operator);
	//   bridgeStream2.applyOperator(operator);
	stream.Of(brigdeStream1, brigdeStream2.SelectChannel("clone")).ApplyOperator(NewVehicleCounter("vehicle counter(shuffle grouping)", 2))

	fmt.Println("This is a streaming job that has one counting operator linked to " + "\n" +
		"two input streams. One operator is configured with default " + "\n" +
		"grouping strategy (shuffle) and the other is configured with fields " + "\n" +
		"grouping strategy. Please enter vehicle types like 'car' and 'truck' in the " + "\n" +
		"input terminal and look at the output")

	starter := job.NewJobStarter(vehicleJob)
	starter.Start()
}
