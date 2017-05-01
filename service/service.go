package service

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/yangyuqian/gateway/types"
	"net/http"
)

// TODO: rewrite url path
// Service represents a group of backend services.
type Service struct {
	Name            string          // service name should be unique, treated as an identity
	Labels          []string        // used for filtering
	Upstream        *types.Upstream // upstream can be nil
	LastEndpointIdx int

	// TODO:
	Authenticators []types.Authenticator // client attempt is accept when there is at least one Authenticator accepts it
	Throttlers     []types.Throttler     // client attempt is limited to min rate
}

// service dispatched by update attributes on request object, which
// can be used in Director of reverse proxy in Go.
// First it lookup a host/path from upstream, which works as a load balancer.
func (svc *Service) Dispatch(r *http.Request) (err error) {
	if r == nil {
		return errors.New("can not dispatch on nil request")
	}

	// 1. balance the upstream endpoints in round robin, set LastEndpointIdx
	// 2. choose next endpoint by default
	ed, err := svc.findEndpoint(r)
	if err != nil {
		return err
	}

	if r.URL == nil {
		return errors.New("can not dispatch on nil URL")
	}

	r.URL.Scheme = ed.Scheme
	r.URL.Host = fmt.Sprintf("%s:%d", ed.Host, ed.Port)

	return
}

// find a specific endpoint
func (svc *Service) findEndpoint(r *http.Request) (ed *types.Endpoint, err error) {
	if svc.Upstream == nil {
		return nil, errors.New("can not determine endpoint on nil upstream")
	}

	if len(svc.Upstream.Endpoints) == 0 {
		return nil, errors.New("can not determine endpoint on empty endpoints")
	}

	svc.LastEndpointIdx = (svc.LastEndpointIdx + 1) % len(svc.Upstream.Endpoints)
	return svc.Upstream.Endpoints[svc.LastEndpointIdx], nil
}
