package types

import (
	"net/http"
)

// Authenticator provides a plugable to verify requests.
type Authenticator interface {
	// accepts a request,
	// verify credentials, then returns true, nil if authenticated successfully
	// authentication should be treated failed either it returns false
	// or error is not nil.
	Do(*http.Request) (bool, error)
}
