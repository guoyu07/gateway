package store

import (
	"github.com/yangyuqian/gateway/service"
	"reflect"
	"testing"
)

const (
	storeAddr = ":57731"
)

func TestRegisterValidService(t *testing.T) {
	for _, cs := range []struct {
		name string
		svc  *service.Service
	}{
		{"valid service", &service.Service{Name: "validsvc1", Labels: []string{"lb1"}}},
		{"valid service name with 10 characters", &service.Service{Name: "validsvc23", Labels: []string{"lb1"}}},
		{"valid service without labels", &service.Service{Name: "validsvc3"}},
		{"valid service with 5 labels", &service.Service{Name: "validsvc4", Labels: []string{"lb1", "lb2", "lb3", "lb4", "lb5"}}},
	} {
		t.Run(cs.name, func(t *testing.T) {
			t.Parallel()
			ServiceRegistry := newServiceRegistry()

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
	for _, cs := range []struct {
		name string
		svc  *service.Service
	}{
		{"nil service", nil},
		{"service with empty name", &service.Service{}},
		{"service with more than 10 characters", &service.Service{Name: "invalidserv"}},
		{"service with more than 5 labels", &service.Service{Name: "svc1", Labels: []string{"lb1", "lb2", "lb3", "lb4", "lb5", "lb6"}}},
		{"service with label of more than 5 characters", &service.Service{Name: "svc2", Labels: []string{"lb1234"}}},
	} {
		t.Run(cs.name, func(t *testing.T) {
			t.Parallel()
			ServiceRegistry := newServiceRegistry()

			err := ServiceRegistry.Register(cs.svc)
			if err == nil {
				t.Errorf("[%s] shouldn't register service(%+v)", cs.name, cs.svc)
			}
		})
	}
}

func TestDeleteServiceByName(t *testing.T) {
	for _, cs := range []struct {
		name string
		svc  *service.Service
	}{
		{"delete service by name", &service.Service{Name: "svc1"}},
	} {
		t.Run(cs.name, func(t *testing.T) {
			t.Parallel()
			ServiceRegistry := newServiceRegistry()

			if err := ServiceRegistry.Register(cs.svc); err != nil {
				t.Errorf("can not register service(%+v) due to(%+v)", cs.svc, err)
			}

			if err := ServiceRegistry.DeleteByName(cs.svc.Name); err != nil {
				t.Errorf("can not delete service by name due to(%+v)", err)
			}
		})
	}
}

func TestServiceFindByLabel(t *testing.T) {
	for _, cs := range []struct {
		name     string
		registry map[string]*service.Service
		label    string
		res      []*service.Service
		wantErr  bool
	}{
		{"find by label", map[string]*service.Service{
			"svc1": &service.Service{Name: "svc1", Labels: []string{"lb1", "lb2"}},
			"svc2": &service.Service{Name: "svc2", Labels: []string{"lb2", "lb3"}},
			"svc3": &service.Service{Name: "svc3", Labels: []string{"lb1"}},
		}, "lb2", []*service.Service{
			&service.Service{Name: "svc1", Labels: []string{"lb1", "lb2"}},
			&service.Service{Name: "svc2", Labels: []string{"lb2", "lb3"}}}, false},
	} {
		t.Parallel()
		ServiceRegistry := newServiceRegistry()

		ServiceRegistry.mutex.Lock()
		ServiceRegistry.registry = cs.registry
		ServiceRegistry.mutex.Unlock()

		svcs, err := ServiceRegistry.FindByLabel(cs.label)
		if !reflect.DeepEqual(svcs, cs.res) {
			t.Errorf("find services(%+v) by label(%s), expected(%+v)", svcs, cs.label, cs.res)
		}

		if err != nil {
			t.Errorf("can not find by label due to(%+v)", err)
		}
	}
}

func TestServiceDeleteByLabel(t *testing.T) {
	for _, cs := range []struct {
		name     string
		registry map[string]*service.Service
		label    string
		res      []*service.Service
		wantErr  bool
	}{
		{"find by label", map[string]*service.Service{
			"svc1": &service.Service{Name: "svc1", Labels: []string{"lb1", "lb2"}},
			"svc2": &service.Service{Name: "svc2", Labels: []string{"lb2", "lb3"}},
			"svc3": &service.Service{Name: "svc3", Labels: []string{"lb1"}},
		}, "lb2", []*service.Service{
			&service.Service{Name: "svc1", Labels: []string{"lb1", "lb2"}},
			&service.Service{Name: "svc2", Labels: []string{"lb2", "lb3"}}}, false},
	} {
		t.Parallel()
		ServiceRegistry := newServiceRegistry()

		ServiceRegistry.mutex.Lock()
		ServiceRegistry.registry = cs.registry
		ServiceRegistry.mutex.Unlock()

		err := ServiceRegistry.DeleteByLabel(cs.label)

		if err != nil {
			t.Errorf("can not delete by label due to(%+v)", err)
		}
	}
}

func TestRegisterServiceRPC(t *testing.T) {
	s, err := NewStoreService(storeAddr)
	if err != nil {
		t.Errorf("can not start store service due to(%+v)", err)
		t.FailNow()
	}

	go s.Serve()

	t.Run("register valid service via RPC", func(t *testing.T) {
		t.Parallel()

		name := "svcrpc1"
		cli, _ := NewStoreClient(storeAddr)
		if err := cli.RegisterService(&service.Service{Name: name}); err != nil {
			t.Errorf("can not register service via rpc due to(%+v)", err)
		}

		ServiceRegistry.mutex.Lock()
		defer ServiceRegistry.mutex.Unlock()

		if _, ok := ServiceRegistry.registry[name]; !ok {
			t.Errorf("service(%s) not registered in registry without any error", name)
		}
	})
}
