package ucs

import (
	"errors"
	"sync"

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

	if handlers, ok := ed.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
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
