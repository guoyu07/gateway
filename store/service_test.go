package store

import (
	"github.com/yangyuqian/gateway/service"
	"reflect"
	"testing"
)

func TestRegisterValidService(t *testing.T) {
	t.Parallel()

	for _, cs := range []struct {
		name string
		svc  *service.Service
	}{
		{"valid service", &service.Service{Name: "svc1", Labels: []string{"lb1"}}},
		{"valid service name with 10 characters", &service.Service{Name: "validservi", Labels: []string{"lb1"}}},
		{"valid service without labels", &service.Service{Name: "svc1"}},
		{"valid service with 5 labels", &service.Service{Name: "svc1", Labels: []string{"lb1", "lb2", "lb3", "lb4", "lb5"}}},
	} {
		t.Run(cs.name, func(t *testing.T) {
			err := ServiceRegistry.Register(cs.svc)
			if err != nil {
				t.Error("can not register service(%+v) due to(%+v)", cs.svc, err)
			}

			ServiceRegistry.mutex.Lock()
			cached, ok := ServiceRegistry.registry[cs.svc.Name]
			if !ok {
				t.Errorf("service is not registered correctly")
			}

			if !reflect.DeepEqual(cached, cs.svc) {
				t.Errorf("registered service(%+v), expected(%+v)", cached, cs.svc)
			}
			ServiceRegistry.mutex.Unlock()
		})
	}
}

func TestRegisterInvalidService(t *testing.T) {
	t.Parallel()

	for _, cs := range []struct {
		name string
		svc  *service.Service
	}{
		{"nil service", nil},
		{"service with empty name", &service.Service{}},
		{"service with more than 10 characters", &service.Service{Name: "invalidserv"}},
		{"service with more than 5 labels", &service.Service{Name: "svc1", Labels: []string{"lb1", "lb2", "lb3", "lb4", "lb5", "lb6"}}},
		{"service with label of more than 5 characters ", &service.Service{Name: "svc2", Labels: []string{"lb1234"}}},
	} {
		t.Run(cs.name, func(t *testing.T) {
			err := ServiceRegistry.Register(cs.svc)
			if err == nil {
				t.Errorf("[%s] shouldn't register service(%+v)", cs.name, cs.svc)
			}
		})
	}
}
