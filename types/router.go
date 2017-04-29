package types

import (
	"net/http"
)

// Router represents a service dispather
type Router interface {
	// dispath incoming request to upstream services
	Dispatch(*http.Request) error
}
