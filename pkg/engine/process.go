package engine

/**
 * This is the base interface of all processes (executors). When a process is started,
 * a new thread is created to call the runOnce() function of the derived class.
 * Each process also have an incoming event queue and an outgoing event queue.
 */
type Process interface {
	Process()

	// Start the process.
	Start()

	// Run process once.return true if the thread should continue; false if the thread should exist.
	runOnce() bool
}

// This struct is used to Inherited by a specific operator
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