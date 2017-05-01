package router

import (
	"github.com/pkg/errors"
	"net/http"
	"regexp"
)

// Matchers are registered in chain, and be treated successful when
// all matchers returns true
type Matcher interface {
	Match(*http.Request) error
}

// Matchers are a group of matchers
type Matchers []Matcher

// add a matcher into the matchers chain
func (ms Matchers) Add(m Matcher) (err error) {
	ms = append(ms, m)
	return
}

// match against incoming request
func (ms Matchers) Match(r *http.Request) (err error) {
	for _, m := range ms {
		if merr := m.Match(r); merr != nil {
			return merr
		}
	}

	return
}

// HostMatcher ensures request.Host == host
func NewHostMatcher(host string) (m *hostMatcher, err error) {
	return &hostMatcher{host: host}, nil
}

type hostMatcher struct {
	host string
}

func (hm *hostMatcher) Match(r *http.Request) (err error) {
	if r == nil {
		return errors.New("can not match on nil request")
	}

	if r.Host == hm.host {
		return nil
	}

	if r.Header != nil && r.Header.Get("Host") == hm.host {
		return nil
	}

	if r.URL != nil && r.URL.Host == hm.host {
		return nil
	}

	return errors.Errorf("host(%s) not match on request", hm.host)
}

// path matcher ensures incoming path matches a regex
func NewPathMatcher(regex string) (pm *pathMatcher, err error) {
	return &pathMatcher{regex: regex}, nil
}

type pathMatcher struct {
	regex string
}

func (pm *pathMatcher) Match(r *http.Request) (err error) {
	matched, err := regexp.MatchString(pm.regex, r.URL.Path)
	if err != nil {
		return errors.Wrap(err, "can not match path")
	}

	if !matched {
		return errors.New("path not matched")
	}

	return
}
