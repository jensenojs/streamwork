package stream

import "streamwork/pkg/engine"

func Of(ss ...Stream) *Streams {
	a := make([]Stream, len(ss))
	for i, s := range ss {
		a[i] = s
	}
	return &Streams{
		array: a,
	}
}

func (ss *Streams) ApplyOperator(op engine.Operator) (engine.Stream, error) {
	for _, s := range ss.array {
		s.ApplyOperator(op)
	}
	return op.GetOutgoingStream(), nil
}
