package stream

import "streamwork/pkg/engine"

func NewWindowedStream(bs *Stream, st engine.WindowStrategy) *WindowedStream {
	return &WindowedStream{
		basestream:        bs,
		windowingStrategy: st,
	}
}

func (w *WindowedStream) ApplyOperator(op engine.WindowOperator) (engine.Stream, error) {
	return w.basestream.applyWindowOperator(op, w.windowingStrategy)
}
