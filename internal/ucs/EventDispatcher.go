package ucs

import (
	"errors"

	"github.com/daividpaulo/EventDispatcherModule/internal/interfaces"
)

type (
	EventDispatcher struct {
		handlers map[string][]interfaces.EventHandlerInterface
	}
)

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]interfaces.EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler interfaces.EventHandlerInterface) error {

	if _, ok := ed.handlers[eventName]; !ok {
		ed.handlers[eventName] = make([]interfaces.EventHandlerInterface, 0)
	}

	if ed.Has(eventName, handler) {
		return errors.New("handler already registered")
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)

	return nil
}

func (ed *EventDispatcher) Dispatch(event interfaces.EventInterface) error {

	if _, ok := ed.handlers[event.GetName()]; !ok {
		return errors.New("event not registered")
	}

	for _, handler := range ed.handlers[event.GetName()] {
		if err := handler.Handle(event); err != nil {
			return err
		}
	}

	return nil
}

func (ed *EventDispatcher) Has(eventName string, handler interfaces.EventHandlerInterface) bool {

	if _, ok := ed.handlers[eventName]; !ok {
		return false
	}

	for _, h := range ed.handlers[eventName] {
		if h == handler {
			return true
		}
	}

	return false
}

func (ed *EventDispatcher) Clear() {
	ed.handlers = make(map[string][]interfaces.EventHandlerInterface)
}
