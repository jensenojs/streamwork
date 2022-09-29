package engine

import "streamwork/pkg/api"

/**
 * This is the base class of all processes (executors). When a process is started,
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

/**
 * The base class for executors of source and operator.
 */
type ComponentExecutor interface {
	GetComponent() api.Component

	SetIncomingQueue(i EventQueue)

	SetOutgoingQueue(i EventQueue)
}
