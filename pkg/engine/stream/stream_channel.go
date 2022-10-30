package stream

import "streamwork/pkg/engine"

func NewStreamChannel(s *Stream, ch engine.Channel) *StreamChannel {
	return &StreamChannel{
		basestream: s,
		channel:    ch,
	}
}

func (sc *StreamChannel) ApplyOperator(op engine.Operator) (engine.Stream, error) {
	return sc.basestream.applyOperator(sc.channel, op)
}
