package api

import 	"errors"


type Stream struct {
	operatorSet map[Operator]bool
}

func NewStream() *Stream {
	return &Stream{
		operatorSet: make(map[Operator]bool),
	}
}

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