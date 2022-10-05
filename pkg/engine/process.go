package engine

/**
 * This is the base interface of all processes (executors). When a process is started,
 * a new thread is created to call the runOnce() function of the derived class.
 * Each process also have an incoming event queue and an outgoing event queue.
 */
type Process interface {
	newProcess()

	// Start the process.
	Start()

	// Run process once.return true if the thread should continue; false if the thread should exist.
	runOnce() bool
}