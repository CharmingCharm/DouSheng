package service

import (
	"context"

	"github.com/CharmingCharm/DouSheng/kitex_gen/base"
	"github.com/CharmingCharm/DouSheng/kitex_gen/video"
)

type LoadVideosService struct {
	ctx context.Context
}

// NewLoadVideosService new LoadVideosService
func NewLoadVideosService(ctx context.Context) *LoadVideosService {
	return &LoadVideosService{ctx: ctx}
}

// CreateUser create user info.
func (s *LoadVideosService) LoadVideos(req *video.LoadVideosRequest) ([]*base.Video, int64, error) {
	// videoList := db.GetVideoListInOrder

	// for v in videoList { rpc.user.GetUserInfo }
	return nil, -1, nil
}
