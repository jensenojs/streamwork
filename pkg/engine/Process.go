package engine


type process struct {
	fn func() 
}

func (p *process) Process() {
	p.fn = func() {
		go func() {
			for {
				p.runOnce()
			}
		}()
	}
}

func (p *process) Start() {
	go p.fn() 
}

func (p *process) runOnce() bool {
	panic("Need to implement runOnce")
}