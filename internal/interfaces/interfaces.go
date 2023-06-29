package interfaces

import "time"

type (
	EventInterface interface {
		GetName() string
		GetDateTime() time.Time
		GetPayload() interface{}
	}

	EventHandlerInterface interface {
		Handle(event EventInterface) error
	}

	EventDispatcherInterface interface {
		Register(eventName string, handler EventHandlerInterface) error
		Dispatch(event EventInterface) error
		Has(eventName string, handler EventHandlerInterface) bool
		Clear()
	}
)
