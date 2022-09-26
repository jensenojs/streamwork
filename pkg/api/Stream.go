package api

import "errors"

/**
 * The Stream class represents a data stream coming out of a component.
 * Operators with the correct type can be applied to this stream.
 * Example:
 *   Job job = new Job("my_job");
 *   job.addSource(mySource)
 *      .applyOperator(myOperator);
 */
type Stream struct {
	operatorSet map[Operator]bool
}

func NewStream() *Stream {
	return &Stream{
		operatorSet: make(map[Operator]bool),
	}
}

/**
 * Apply an operator to this stream.
 * @param operator The operator to be connected to the current stream
 * @return The outgoing stream of the operator.
 */
func (s *Stream) ApplyOperator(op Operator) (stream *Stream, err error) {
	if _, ok := s.operatorSet[op]; ok {
		err = errors.New("Operator " + op.GetName() + " already exists")
		return
	}
	s.operatorSet[op] = true
	return op.GetOutgoingStream(), nil
}

func (s *Stream) GetAppliedOperators() map[Operator]bool {
	return s.operatorSet
}
