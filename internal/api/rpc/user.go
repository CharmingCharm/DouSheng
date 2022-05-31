package rpc

import (
	"context"

	"github.com/CharmingCharm/DouSheng/kitex_gen/user"
	"github.com/CharmingCharm/DouSheng/kitex_gen/user/userservice"
	"github.com/CharmingCharm/DouSheng/pkg/constants"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
		// client.WithMiddleware(middleware.CommonMiddleware),
		// client.WithInstanceMW(middleware.ClientMiddleware),
		// client.WithMuxConnection(1),                     // mux
		// client.WithRPCTimeout(3*time.Second),            // rpc timeout
		// client.WithConnectTimeout(50*time.Millisecond),  // conn timeout、
		client.WithSuite(trace.NewDefaultClientSuite()), // tracer
		client.WithResolver(r),                          // resolver
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// CreateUser create user info
func CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	resp, err := userClient.CreateUser(ctx, req)
	return resp, err
}
