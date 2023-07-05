package interfaces

import (
	"sync"
	"time"
)

type (
	EventInterface interface {
		GetName() string
		GetDateTime() time.Time
		GetPayload() interface{}
	}

	EventHandlerInterface interface {
		Handle(event EventInterface, wg *sync.WaitGroup) error
	}

	EventDispatcherInterface interface {
		Register(eventName string, handler EventHandlerInterface) error
		Dispatch(event EventInterface) error
		Has(eventName string, handler EventHandlerInterface) bool
		Clear()
	}
)
