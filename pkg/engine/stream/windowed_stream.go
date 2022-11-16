package stream

import "streamwork/pkg/engine"

func NewWindowedStream(bs *Stream, st engine.GroupStrategy) *WindowedStream {
	return &WindowedStream{
		basestream:        bs,
		windowingStrategy: st,
	}
}

func (w *WindowedStream) ApplyOperator(op engine.Operator) (engine.Stream, error) {
	return w.basestream.ApplyOperator(op)
}
