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
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := actionservice.NewClient(
		constants.ActionServiceName,
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

// func GetCommentList(ctx context.Context, req *action.GetCommentListRequest) (*action.GetCommentListResponse, error) {
// 	// TODO
// 	resp, err := actionClient.GetCommentList(ctx, req)
// 	return resp, err
// }

// func GetFavoriteVideos(ctx context.Context, req *action.GetFavoriteVideosRequest) (*action.GetFavoriteVideosResponse, error) {
// 	// TODO
// 	resp, err := actionClient.GetFavoriteVideos(ctx, req)
// 	return resp, err
// }

// func GetUserFollowerList(ctx context.Context, req *action.GetUserFollowerListRequest) (*action.GetUserFollowerListResponse, error) {
// 	// TODO
// 	resp, err := actionClient.GetUserFollowerList(ctx, req)
// 	return resp, err
// }

// func GetUserFollowList(ctx context.Context, req *action.GetUserFollowListRequest) (*action.GetUserFollowListResponse, error) {
// 	// TODO
// 	resp, err := actionClient.GetUserFollowList(ctx, req)
// 	return resp, err
// }

// func UpdateComment(ctx context.Context, req *action.UpdateCommentRequest) (*action.UpdateCommentResponse, error) {
// 	// TODO
// 	resp, err := actionClient.UpdateComment(ctx, req)
// 	return resp, err
// }

// func UpdateFavorite(ctx context.Context, req *action.UpdateFavoriteRequest) (*action.UpdateFavoriteResponse, error) {
// 	// TODO
// 	resp, err := actionClient.UpdateFavorite(ctx, req)
// 	return resp, err
// }

// func UpdateRelationship(ctx context.Context, req *action.UpdateRelationshipRequest) (*action.UpdateRelationshipResponse, error) {
// 	// TODO
// 	resp, err := actionClient.UpdateRelationship(ctx, req)
// 	return resp, err
// }
