package service

import (
	"context"

	"github.com/CharmingCharm/DouSheng/kitex_gen/base"
	"github.com/CharmingCharm/DouSheng/kitex_gen/video"
)

type GetPublishedVideosService struct {
	ctx context.Context
}

// NewGetPublishedVideosService new GetPublishedVideosService
func NewGetPublishedVideosService(ctx context.Context) *GetPublishedVideosService {
	return &GetPublishedVideosService{ctx: ctx}
}

// CreateUser create user info.
func (s *GetPublishedVideosService) GetPublishedVideos(req *video.GetPublishedVideosRequest) ([]*base.Video, error) {
	// videoList := db.GetVideoListByAuthorId

	// for v in videoList { rpc.user.GetUserInfo }
	return nil, nil
}
