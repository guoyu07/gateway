package router

import (
	"net/http"
	"testing"
)

func TestRouterMatch(t *testing.T) {
	for _, cs := range []struct {
		name    string
		router  *Router
		req     *http.Request
		wantErr bool
	}{
		{
			"host not matched",
			&Router{},
			&http.Request{},
			true,
		},
		{
			"host matched",
			&Router{Matchers: Matchers{&hostMatcher{host: "www.example.com"}}, Service: "svc1"},
			&http.Request{Host: "www.example.com"},
			false,
		},
	} {
		t.Run(cs.name, func(t *testing.T) {
			if _, err := cs.router.Match(cs.req); (err != nil) != cs.wantErr {
				t.Errorf("match failed due to(%+v)", err)
			}
		})
	}
}
