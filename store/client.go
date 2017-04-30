package store

import (
	"github.com/pkg/errors"
	protobuf "github.com/yangyuqian/gateway/protobuf"
	"github.com/yangyuqian/gateway/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func NewStoreClient(addr string) (cli *storeClient, err error) {
	cli = &storeClient{addr: addr}

	return
}

type storeClient struct {
	addr string
}

func (c *storeClient) RegisterService(svc *service.Service) (err error) {
	conn, err := grpc.Dial(c.addr, grpc.WithInsecure())
	if err != nil {
		return errors.Wrapf(err, "can not connect to(%s)", c.addr)
	}
	defer conn.Close()

	cli := protobuf.NewServiceRegistryClient(conn)
	reply, err := cli.RegisterService(context.Background(), &protobuf.ServiceRequest{Name: svc.Name, Labels: svc.Labels})
	if err != nil {
		return errors.Wrapf(err, "can not register service(%+v)", svc)
	}

	if !reply.Ok {
		return errors.Errorf("can not register service(%+v) due to(%s)", svc, reply.GetResult())
	}

	return
}
