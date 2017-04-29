package types

// Endpoint represents the a interface
// for a upstream which is used to do a load balancing
type Endpoint struct {
	Protocol string // either TCP or UDP
	Host     string // endpoint host without NAT
	Port     int    // endpoint port
}

// Upstream represents a group of endpoints
// with load balancing in front of them.
type Upstream struct {
	Name      string      // name should be unique
	Labels    []string    // labels used to filter
	Endpoints []*Endpoint // optional endpoints
}
