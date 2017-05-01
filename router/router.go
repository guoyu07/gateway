package router

// Router is a rule to determine a matcher request to be
// dispatched to specific service.
type Router struct {
	// matchers determines the request, matchers can be empty, which means
	// all request will be accepted by this rule
	Matchers Matchers
	// destination service when request is matched by all matchers
	Service string
}
