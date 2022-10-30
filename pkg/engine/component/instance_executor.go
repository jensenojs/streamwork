package component

import (
	"streamwork/pkg/engine"
)

// TakeIncomingEvent is a helper function to receive events
func (i *InstanceExecutorImpl) TakeIncomingEvent() engine.Event {
	return i.Incoming.Take()
}

// SendOutgoingEvent is a helper function to send events to all downstreams.
func (i *InstanceExecutorImpl) SendOutgoingEvent() {
	for _, ch := range i.EventCollector.GetRegisteredChannels() {
		for _, out := range i.EventCollector.GetEventList(ch) {
			for _, q := range i.OutgoingMap[ch] {
				q.Send(out)
			}
		}
	}
}

// =================================================================
// implement for Process

func (i *InstanceExecutorImpl) NewProcess() {
	i.FnWrapper = func() {
		for {
			if ok := i.Fn(); ok != true {
				break
			}
		}
	}
}

func (i *InstanceExecutorImpl) Start() {
	go i.FnWrapper()
}

func (i *InstanceExecutorImpl) RunOnce() bool {
	panic("Need specific implementation")
}

// SetRunOnce is a helper function to set RunOnce during operator / source executor impl init.
func (i *InstanceExecutorImpl) SetRunOnce(RunOnce func() bool) {
	i.Fn = RunOnce
}

// =================================================================
// implement for InstanceExecutor

func (i *InstanceExecutorImpl) SetIncoming(in engine.EventQueue) {
	i.Incoming = in
}

func (i *InstanceExecutorImpl) AddOutgoing(ch engine.Channel, out engine.EventQueue) {
	if c, ok := i.OutgoingMap[ch]; ok {
		c = append(c, out)
	} else {
		l := make([]engine.EventQueue, 0)
		l[0] = out
		i.OutgoingMap[ch] = l
	}
}

func (i *InstanceExecutorImpl) RegisterChannel(ch engine.Channel) {
	i.EventCollector.SetRegisterChannel(ch)
}