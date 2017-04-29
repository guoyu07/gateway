package event

// EventType represents type of events, such as creating/removing a upstream
type EventType int

const (
	// add a service
	ADD_SERVICE EventType = iota
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
