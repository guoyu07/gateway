package event

import (
	"github.com/yangyuqian/gateway/service"
)

// EventType represents type of events, such as creating/removing a upstream
type EventType int

const (
	// join/leave cluster
	JOIN EventType = iota
	LEAVE

	// add a service
	ADD_SERVICE
	// delete a service
	DEL_SERVICE
	// add upstream to service
	ADD_UPSTREAM
	// delete upstream from service
	DEL_UPSTREAM
	// add router
	ADD_ROUTER
	// delete router
	DEL_ROUTER
)

// Event represents a message to ask the cluster to co-operate
type Event struct {
	// type of the event
	Type EventType
	// payload in bytes
	Payload []byte
}

var (
	// read/write channel
	regSvc = make(chan service.Service, 10)
	// expose write-only channel
	RegSvcCh chan<- service.Service = regSvc
)
