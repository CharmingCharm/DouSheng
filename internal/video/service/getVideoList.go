package service

import (
	"context"

	"github.com/CharmingCharm/DouSheng/kitex_gen/base"
	"github.com/CharmingCharm/DouSheng/kitex_gen/video"
)

type GetVideoListService struct {
	ctx context.Context
}

// NewGetVideoListService new GetVideoListService
func NewGetVideoListService(ctx context.Context) *GetVideoListService {
	return &GetVideoListService{ctx: ctx}
}

// CreateUser create user info.
func (s *GetVideoListService) GetVideoList(req *video.GetVideoListRequest) ([]*base.Video, error) {
	// videoList := db.GetVideoListByIds

	// for v in videoList { rpc.user.GetUserInfo }
	return nil, nil
}
