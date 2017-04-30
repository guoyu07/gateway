package service

import (
	"bytes"
	"github.com/hashicorp/go-msgpack/codec"
	"github.com/yangyuqian/gateway/types"
)

// Service represents a group of backend services.
type Service struct {
	Name           string                // service name should be unique, treated as an identity
	Labels         []string              // used for filtering
	Upstream       *types.Upstream       // upstream can be nil
	Authenticators []types.Authenticator // client attempt is accept when there is at least one Authenticator accepts it
	Throttlers     []types.Throttler     // client attempt is limited to min rate
}
