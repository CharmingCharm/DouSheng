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
		client.WithSuite(trace.NewDefaultClientSuite()), // tracer
		client.WithResolver(r),                          // resolver
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	resp, err := userClient.GetUserInfo(ctx, req)
	return resp, err
}

func UpdateUserFollow(ctx context.Context, req *user.UpdateUserFollowRequest) (*user.UpdateUserFollowResponse, error) {
	resp, err := userClient.UpdateUserFollow(ctx, req)
	return resp, err
}
