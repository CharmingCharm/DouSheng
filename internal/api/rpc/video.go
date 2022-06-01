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
	videoClient = c
}

func LoadVideos(ctx context.Context, req *video.LoadVideosRequest) (*video.LoadVideosResponse, error) {
	// TODO
	// var loadVideosResponse *video.LoadVideosResponse
	// loadVideosResponse, err = videoClient.LoadVideos(ctx, req)
	//
	// return loadVideosResponse, err
	return nil, nil
}
