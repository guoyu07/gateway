package event

import (
	"github.com/hashicorp/serf/serf"
)

// new broadcast, it's necessary to have independent addrs:
// bindAddr and advertiseAddr here, bindAddr for binding address on the *host*
// and advertiseAddr for exposed address to other nodes.
// This will be useful when the agent running inside a NAT or container.
func NewBroadcast(name, bindAddr, advertiseAddr string) (b *broadcast) {
	return
}

// broadcast broadcasts messages to the cluster
type broadcast struct {
	serfInstance *serf.Serf
}
