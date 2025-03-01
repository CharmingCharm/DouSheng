package rpc

import (
	"context"

	"github.com/CharmingCharm/DouSheng/kitex_gen/action"
	"github.com/CharmingCharm/DouSheng/kitex_gen/action/actionservice"
	"github.com/CharmingCharm/DouSheng/pkg/constants"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var actionClient actionservice.Client

func initActionRpc() {
	// Register for an action service client
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	// Create the client
	c, err := actionservice.NewClient(
		constants.ActionServiceName,
		client.WithSuite(trace.NewDefaultClientSuite()), // tracer
		client.WithResolver(r),                          // resolver
	)
	if err != nil {
		panic(err)
	}
	actionClient = c
}

func CheckRelation(ctx context.Context, req *action.CheckRelationRequest) (*action.CheckRelationResponse, error) {
	// TODO
	resp, err := actionClient.CheckRelation(ctx, req)
	return resp, err
}

func CheckFavorite(ctx context.Context, req *action.CheckFavoriteRequest) (*action.CheckFavoriteResponse, error) {
	// TODO
	resp, err := actionClient.CheckFavorite(ctx, req)
	return resp, err
}
