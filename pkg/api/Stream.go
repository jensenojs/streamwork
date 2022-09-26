package api

type void struct{}

type Stream struct {
	operatorSet map[Operator]void
}

func NewStream() Stream {
	panic("unimplemented")
}