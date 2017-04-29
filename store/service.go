package store

import (
	"github.com/pkg/errors"
	"github.com/yangyuqian/gateway/service"
	"strings"
	"sync"
)

var ServiceRegistry = newServiceRegistry()

func newServiceRegistry() *serviceRegistry {
	return &serviceRegistry{registry: make(map[string]*service.Service)}
}

// serviceRegistry is a in-memory datastore
// for registered services
type serviceRegistry struct {
	// locker to protect the service store
	mutex sync.Mutex
	// registry of services, the key is service name
	registry map[string]*service.Service
}

func (r *serviceRegistry) Register(svc *service.Service) (err error) {
	if svc == nil {
		return errors.New("can not register nil service")
	}

	svc.Name = strings.Trim(svc.Name, " ")
	if svc.Name == "" {
		return errors.New("can not register service with empty name or name with all whitespaces")
	}

	if len(svc.Name) > 10 {
		return errors.Errorf("invalid service name(%s), service names can have at most 10 characters", svc.Name)
	}

	if len(svc.Labels) > 5 {
		return errors.Errorf("can not register service(%s) with more than 5 labels(%+v)", svc.Name, svc.Labels)
	}

	for _, lb := range svc.Labels {
		if len(lb) > 5 {
			return errors.Errorf("invalid service label(%s)", lb)
		}

		if strings.Trim(lb, " ") == "" {
			return errors.New("service label not be empty text or all whitespaces")
		}
	}

	r.mutex.Lock()

	if r.registry == nil {
		return errors.Errorf("can not register service(%s) due to registry is not initialized properly", svc.Name)
	}

	r.registry[svc.Name] = svc

	r.mutex.Unlock()

	return
}
