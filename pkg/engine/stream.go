package engine

import (
	"errors"
)

/**
 * The Stream class represents a data stream coming out of a component.
 * Operators with the correct type can be applied to this stream.
 * Example:
 *   Job job = new Job("my_job");
 *   job.addSource(mySource)
 *      .applyOperator(myOperator);
 */
type Stream struct {
	OperatorSet map[Operator]Void
}

// helper function to init a stream
func NewStream() *Stream {
	return &Stream{
		OperatorSet: make(map[Operator]Void),
	}
}

// ApplyOperator Apply an operator to this stream, return the outgoing stream of the operator.
func (s *Stream) ApplyOperator(op Operator) (stream *Stream, err error) {
	if _, ok := s.OperatorSet[op]; ok {
		err = errors.New("Operator " + op.GetName() + " already exists")
		return
	}
	s.OperatorSet[op] = Member
	return op.GetOutgoingStream(), nil
}

func (s *Stream) GetAppliedOperators() map[Operator]Void {
	return s.OperatorSet
}
