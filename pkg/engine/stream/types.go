// The Stream class represents a data stream coming out of a component.
// Operators with the correct type can be applied to this stream.
package stream

import "streamwork/pkg/engine"

// Example:
//
//	Job job = new Job("my");
//	job.addSource(mySource)
//	   .applyOperator(myOperator);
type Stream struct {
	Channel     engine.Channel
	OperatorMap map[engine.Channel]map[engine.Operator]engine.Void
}

// Streams will be used in stream-graph, like split and merge
type Streams struct {
	array []engine.Stream
}

// Example:
//
//	Job job = new Job("my");
//	job.addSource(mySource)
//	   .selectChannel("my_channel")
//	   .applyOperator(myOperator);
type StreamChannel struct {
	channel    engine.Channel
	basestream *Stream
}

// Example:
//
//	Job job = new Job("my_job");
//	job.addSource(mySource)
//	   .withWindowing(new FixedTimeWindowingStrategy(1000, 1000))
//	   .applyOperator(myOperator);
type WindowedStream struct {
	windowingStrategy engine.WindowStrategy
	basestream        *Stream
}
