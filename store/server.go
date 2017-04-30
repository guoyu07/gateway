package store

import (
	"github.com/pkg/errors"
	protobuf "github.com/yangyuqian/gateway/protobuf"
	"github.com/yangyuqian/gateway/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

// create and run a rpc service for datastore
func NewStoreService(addr string) (s *storeService, err error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, errors.Wrapf(err, "can not listen(%s)", addr)
	}

	server := grpc.NewServer()
	s = &storeService{listener: lis, server: server}

	protobuf.RegisterServiceRegistryServer(server, s)
	reflection.Register(server)
	return
}

// StoreService contains functions serving rpc interfaces to
// interact to the built-in datastore
type storeService struct {
	listener net.Listener
	server   *grpc.Server
}

func (s *storeService) Serve() (err error) {
	if s.listener == nil {
		return errors.New("can not serve on nil listener")
	}

	if s.server == nil {
		return errors.New("can not serve on nil server")
	}

	if err := s.server.Serve(s.listener); err != nil {
		return errors.Wrap(err, "error serving store service")
	}

	return
}

func (s *storeService) RegisterService(ctx context.Context, req *protobuf.ServiceRequest) (resp *protobuf.ServiceReply, err error) {
	if regerr := ServiceRegistry.Register(&service.Service{Name: req.GetName(), Labels: req.GetLabels()}); regerr != nil {
		resp = &protobuf.ServiceReply{Ok: false, Result: regerr.Error()}
		return
	}
	return &protobuf.ServiceReply{Ok: true, Result: "register service successfully"}, nil
}
