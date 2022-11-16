# streamwork

study from [GrokkingStreamingSystems](https://github.com/nwangtw/GrokkingStreamingSystems)
, and implementing in go way.

## brief introduction

For reading the code, the first recommended version is v0.4.

* Previous versions had unreasonable package divisions (v0.1, v0.2).
* Had some practically useless code that was not removed in the refactoring process(v0.3).
* v0.4 actually still has some unreasonable places, but it's readability should be even better than the original version, [GrokkingStreamingSystems](https://github.com/nwangtw/GrokkingStreamingSystems)

## What's in v0.4 now

* support custom define job, as long as you implement your own source and operator executors.
* suport run parallelly
  * support round-robin strategy and filed grouping strategy
* support stream-graph
* still processing each element individally time (i.e. no window function)
* at-most-once(i.e. no guarantees of being successfully processed)

## Design

Design idea is very Straightforward. Operator and Source are both Component in Stream. You can see it in pkg/engine/types.go

### Component

Based on above, Operator interface requires to implement `Apply(Event, EventCollector) error`, This function needs to be implemented by users.
Similarly, Source interface requires to implement `GetEvents(string, EventCollector)`. The first argument actually is event, just encode as string, and will put the event into EventCollector to dispatch downstream operator.

The stream consists of only these two components —— the source and the operator. Component implements the interface and is inherited by source.Source and operator.Operator, which implement the necessary methods for reuse by all operators. On top of that, the user only needs to care about how their business logic is implemented.

### ComponentExecutor

After defining the operator, the question to discuss is how streamwork will run the operator, which involves the executor.The executor design is similar to the above, with a component executor that will be inherited by the source executor and operator executor.

But there are two points different from the component
    1. User does not need to care about this detail, The user only needs to care about how `Apply(Event, EventCollector) error` and `GetEvents(string, EventCollector)` are implemented.
    2. Because of the need to support parallelism, so there may be multiple instance executors to handle event separately.

### Others

* Process interface defines how executors will run
* event collector, event dispatcher and event queue are helper data structures to transport event through the stream. As we needs to support stream-graph, so one event may be dispatched to multiple operators and multiple event may merged in one operator, we need some mechanism to support it.
* TODO.?

## How to define your own job

two steps

1. define your own operator and source
2. uses your operator and source to create a new job and start it

#### How to define your own Operator and Source

Take vehicle count job as example, just inherit source.Source and operator.Operator, and add whatever data structer you need to support your logic in `Apply(Event, EventCollector) error` or `GetEvents(string, EventCollector)`.

```go
type SensorReader struct {
 source.Source
}

type VehicleCounter struct {
 operator.Operator
 counter map[carType]int
}
```

### How to define your own job

Take a look at the below example, you can find more information in pkg/jobs/vehicle_count_job/job_test.go, just use ApplyOperator to connect operator to each others.

```go
 vehicleJob := job.NewJob("vehicle count base test")
 brigdeStream, err := vehicleJob.AddSource(NewSensorReader("sensor-reader"))
 if err != nil {
  panic(err)
 }
 brigdeStream.ApplyOperator(NewVehicleCounter("vehicle counter"))
```

once you start it, you can open another terminal and use `nc localhost 9990` to connect to streamwork

```go
 starter := job.NewJobStarter(vehicleJob)
 starter.Start()
```

## How to get start

1. git clone git@github.com:jensenojs/streamwork.git
2. install go and necessary dependencies
3. `go run pkg/main.go`
4. once you start it, you can open another terminal and use `nc localhost 9990` to connect to streamwork

ps : if you want to debug the job in package jobs, make sure you use debug mode so the output can be shown in debug console.
