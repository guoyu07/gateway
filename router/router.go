package router

import (
	"github.com/pkg/errors"
	"net/http"
)

// Router is a rule to determine a matcher request to be
// dispatched to specific service.
type Router struct {
	// matchers determines the request, matchers can be empty, which means
	// all request will be accepted by this rule
	Matchers Matchers
	// destination service when request is matched by all matchers
	Service string
	// priority of the router
	Priority uint8
}

// determine upstream with request
func (rt *Router) Match(r *http.Request) (svc string, err error) {
	if len(rt.Matchers) == 0 {
		return "", errors.New("can not determine routes with empty matchers")
	}

	if err = rt.Matchers.Match(r); err != nil {
		return "", errors.Wrap(err, "route not matched")
	}

	return rt.Service, nil
}
