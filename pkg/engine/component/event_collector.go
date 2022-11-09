package component

import "streamwork/pkg/engine"

func NewEventCollector() *EventCollector {
	return &EventCollector{
		defaultChannel:   engine.DEFAULT_CHANNEL,
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
	}
}

func (e *EventCollector) GetEventList(ch engine.Channel) []engine.Event {
	return e.List[ch]
}

func (e *EventCollector) Add(ev engine.Event) {
	e.Addto(ev, e.defaultChannel)
}

func (e *EventCollector) Addto(ev engine.Event, ch engine.Channel) {
	// If the channel is registered, add the event to the corresponding list.
	if _, ok := e.RegisterChannels[ch]; !ok {
		panic("unknown channel")
	}

	l := e.List[ch]

	if len(l) == 0 {
		l = make([]engine.Event, 1)
		l[0] = ev
		e.List[ch] = l
	} else {
		l = append(l, ev)
	}
}

func (e *EventCollector) Clear() {
	for k := range e.List {
		e.List[k] = nil
	}
}
