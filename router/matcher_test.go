package router

import (
	"net/http"
	"net/url"
	"testing"
)

func TestMatcher(t *testing.T) {
	for _, cs := range []struct {
		name    string
		matcher Matcher
		req     *http.Request
		wantErr bool
	}{
		{
			"host not matched",
			&hostMatcher{host: "www.example.com"},
			&http.Request{},
			true,
		},
		{
			"host matched",
			&hostMatcher{host: "www.example.com"},
			&http.Request{Host: "www.example.com"},
			false,
		},
		{
			"path matched",
			&pathMatcher{regex: "^/p1/p2$"},
			&http.Request{URL: &url.URL{Path: "/p1/p2"}},
			false,
		},
		{
			"path matched",
			&pathMatcher{regex: "^/p1/p2$"},
			&http.Request{URL: &url.URL{Path: "/p1"}},
			true,
		},
		{
			"path or host not matched",
			Matchers{&pathMatcher{regex: "^/p1/p2$"}, &hostMatcher{host: "www.example.com"}},
			&http.Request{Host: "www.example.com", URL: &url.URL{Path: "/p1/p2"}},
			false,
		},
	} {
		t.Run(cs.name, func(t *testing.T) {
			if err := cs.matcher.Match(cs.req); (err != nil) != cs.wantErr {
				t.Errorf("match failed due to(%+v)", err)
			}
		})
	}
}
