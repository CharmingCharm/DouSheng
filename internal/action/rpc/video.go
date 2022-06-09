package rpc

import (
	"context"

	"github.com/CharmingCharm/DouSheng/kitex_gen/video"
	"github.com/CharmingCharm/DouSheng/kitex_gen/video/videoservice"
	"github.com/CharmingCharm/DouSheng/pkg/constants"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var videoClient videoservice.Client

func initVideoRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := videoservice.NewClient(
		constants.VideoServiceName,
		client.WithSuite(trace.NewDefaultClientSuite()), // tracer
		client.WithResolver(r),                          // resolver
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func GetVideoList(ctx context.Context, req *video.GetVideoListRequest) (*video.GetVideoListResponse, error) {
	// TODO
	resp, err := videoClient.GetVideoList(ctx, req)
	return resp, err
}

func UpdateFavoriteCount(ctx context.Context, req *video.UpdateFavoriteCountRequest) (*video.UpdateFavoriteCountResponse, error) {
	// TODO
	resp, err := videoClient.UpdateFavoriteCount(ctx, req)
	return resp, err
}

func UpdateCommentCount(ctx context.Context, req *video.UpdateCommentCountRequest) (*video.UpdateCommentCountResponse, error) {
	// TODO
	resp, err := videoClient.UpdateCommentCount(ctx, req)
	return resp, err
}
