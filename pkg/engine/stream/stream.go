package stream

import (
	"errors"
	"streamwork/pkg/engine"
)

// helper function to init a stream
func NewStream() *Stream {
	return &Stream{
		Channel:     engine.DEFAULT_CHANNEL,
		OperatorMap: make(map[engine.Channel]map[engine.Operator]engine.Void),
	}
}

// ApplyOperator Apply an operator to this stream, return the outgoing stream of the operator.
func (s *Stream) ApplyOperator(op engine.Operator) (engine.Stream, error) {
	return s.applyOperator(s.Channel, op)
}

// applyOperator Apply an operator to specified channel of this stream, return the outgoing stream of the operator.
func (s *Stream) applyOperator(ch engine.Channel, op engine.Operator) (engine.Stream, error) {
	if set, ok := s.OperatorMap[ch]; ok {
		if _, ok := set[op]; ok {
			return nil, errors.New("Operator " + op.GetName() + " already exists")
		}
		set[op] = engine.Member
	} else {
		//  a new channel
		set := make(map[engine.Operator]engine.Void)
		set[op] = engine.Member
		s.OperatorMap[ch] = set
	}
	return op.GetOutgoingStream(), nil
}

func (s *Stream) selectChannel(ch engine.Channel) *StreamChannel {
	return NewStreamChannel(s, ch)
}

// Get the channels in the stream. Note that the channel set
// is collected from the downstream component's applyOperator() calls.
func (s *Stream) GetChannels() []engine.Channel {
	chs := make([]engine.Channel, len(s.OperatorMap))
	i := 0
	for c := range s.OperatorMap {
		chs[i] = c
		i++
	}
	return chs
}

// Get the collection of operators applied to this stream.
func (s *Stream) GetAppliedOperators(ch engine.Channel) map[engine.Operator]engine.Void {
	return s.OperatorMap[ch]
}
