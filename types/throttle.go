package types

import (
	"net/http"
)

// Throttler represents the object to limit rate of client requests
type Throttler interface {
	// return a reason(string) and nil error if need to be throttled
	// or else it should be treated as accepted
	Do(*http.Request) (string, error)
}
