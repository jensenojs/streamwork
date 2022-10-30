package component

import "streamwork/pkg/engine"

func NewEventCollector() *EventCollector {
	return &EventCollector{
		DEFAULT_CHANNEL:  engine.DEFAULT_CHANNEL,
		List:             make(map[engine.Channel][]engine.Event),
		RegisterChannels: make(map[engine.Channel]engine.Void),
	}
}

func (e *EventCollector) GetRegisteredChannels() []engine.Channel {
	chs := make([]engine.Channel, len(e.RegisterChannels))
	i := 0
	for ch := range e.RegisterChannels {
		chs[i] = ch
		i++
	}
	return chs
}

func (e *EventCollector) SetRegisterChannel(ch engine.Channel) {
	if _, ok := e.RegisterChannels[ch]; !ok {
		e.RegisterChannels[ch] = engine.Member
		e.List[ch] = make([]engine.Event, 0)
	}
}

func (e *EventCollector) GetEventList(ch engine.Channel) []engine.Event {
	return e.List[ch]
}

func (e *EventCollector) Add(ev engine.Event) {
	e.Addto(ev, e.DEFAULT_CHANNEL)
}

func (e *EventCollector) Addto(ev engine.Event, ch engine.Channel) {
	// If the channel is registered, add the event to the corresponding list.
	if l, ok := e.List[ch]; ok {
		l = append(l, ev)
	}
}

func (e *EventCollector) Clear() {
	e.List = nil
	e.RegisterChannels = nil
}
