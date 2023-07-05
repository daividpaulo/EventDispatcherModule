package models

import (
	"log"
	"sync"
	"testing"
	"time"

	"github.com/daividpaulo/EventDispatcherModule/internal/interfaces"
	"github.com/daividpaulo/EventDispatcherModule/internal/ucs"
)

type NodeCreatedEvent struct {
	Name     string
	DateTime time.Time
	NodeID   string
	SiteID   string
}

func (e *NodeCreatedEvent) GetName() string {
	return "node.created"
}

func (e *NodeCreatedEvent) GetDateTime() time.Time {
	return e.DateTime
}

func (e *NodeCreatedEvent) GetPayload() interface{} {
	return *e
}

func NewNodeCreatedEvent(nodeID string, siteID string) *NodeCreatedEvent {
	return &NodeCreatedEvent{
		Name:     "node.created",
		DateTime: time.Now(),
		NodeID:   nodeID,
		SiteID:   siteID,
	}
}

// create a handler
type KrakenSyncHandler struct {
}

func (h *KrakenSyncHandler) Handle(event interfaces.EventInterface, wg *sync.WaitGroup) error {

	nodeCreatedEvent := event.GetPayload().(NodeCreatedEvent)

	log.Printf("KrakenSyncHandler %s", nodeCreatedEvent.NodeID)
	//log.Printf("Node %s created at site %s name %s", nodeCreatedEvent.NodeID, nodeCreatedEvent.Name, nodeCreatedEvent.GetName())

	wg.Done()

	return nil
}

type SyncLcHandler struct {
}

func (h *SyncLcHandler) Handle(event interfaces.EventInterface, wg *sync.WaitGroup) error {

	nodeCreatedEvent := event.GetPayload().(NodeCreatedEvent)

	log.Printf("SyncLcHandler %s", nodeCreatedEvent.NodeID)
	//	log.Printf("Node %s created at site %s name %s", nodeCreatedEvent.NodeID, nodeCreatedEvent.Name, nodeCreatedEvent.GetName())

	wg.Done()

	return nil
}

func TestGeral(t *testing.T) {
	// create a dispatcher
	dispatcher := ucs.NewEventDispatcher()

	// register event handlers
	dispatcher.Register("node.created", &KrakenSyncHandler{})
	dispatcher.Register("node.created", &SyncLcHandler{})

	// dispatch event in side service
	dispatcher.Dispatch(NewNodeCreatedEvent("123", "456"))
}
