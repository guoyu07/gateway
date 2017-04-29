package types

// Service represents a group of backend services.
type Service struct {
	Name           string           // service name should be unique, treated as an identity
	Labels         []string         // used for filtering
	Upstream       *Upstream        // upstream can be nil
	Authenticators []*Authenticator // client attempt is accept when there is at least one Authenticator accepts it
	Throttlers     []*Throttler     // client attempt is limited to min rate
}
