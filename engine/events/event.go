package events

import (
	"fmt"
)

type Event interface {
	Key() string
	Contents() interface{}
	Handled() bool
	Handle()
}

type event struct {
	key      string
	contents interface{}
	handled  bool
}

func NewEvent(key string, contents interface{}) Event {
	return &event{
		key:      key,
		contents: contents,
		handled:  false,
	}
}

func (e *event) Key() string {
	return e.key
}

func (e *event) Contents() interface{} {
	return e.contents
}

func (e *event) Handled() bool {
	return e.handled
}

func (e *event) Handle() {
	e.handled = true
}

type Listener interface {
	OnEvent(event Event)
}

type Publisher interface {
	Listen(sub Listener)
	Broadcast(event Event) error
}

type eventQue struct {
	listeners []Listener
}

func NewPublisher() Publisher {
	return &eventQue{
		listeners: []Listener{},
	}
}

func (que *eventQue) Listen(sub Listener) {
	que.listeners = append(que.listeners, sub)
}

func (que *eventQue) Broadcast(event Event) error {
	if len(que.listeners) == 0 {
		return fmt.Errorf("no event listeners")
	}
	for _, sub := range que.listeners {
		sub.OnEvent(event)
	}
	return nil
}
