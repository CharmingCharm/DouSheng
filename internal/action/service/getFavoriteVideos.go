package service

import (
	"context"

	"github.com/CharmingCharm/DouSheng/internal/action/db"
	"github.com/CharmingCharm/DouSheng/internal/action/rpc"

	"github.com/CharmingCharm/DouSheng/kitex_gen/action"
	"github.com/CharmingCharm/DouSheng/kitex_gen/base"
	"github.com/CharmingCharm/DouSheng/kitex_gen/video"
)

type GetFavoriteVideosService struct {
	ctx context.Context
}

// NewGetFavoriteVideosService new GetFavoriteVideosService
func NewGetFavoriteVideosService(ctx context.Context) *GetFavoriteVideosService {
	return &GetFavoriteVideosService{ctx: ctx}
}

// Get the "Favorite" videos
func (s *GetFavoriteVideosService) GetFavoriteVideos(req *action.GetFavoriteVideosRequest) ([]*base.Video, error) {
	// Get the ids of the video list that is favorite by the user
	vIds, err := db.GetFavoriteVideoIdsByUserId(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	// Use rpc call to fetch video list from video server
	resp, err := rpc.GetVideoList(s.ctx, &video.GetVideoListRequest{
		VideoIds: vIds,
		UserId:   req.MyId,
	})
	if err != nil {
		return nil, err
	}

	return resp.Videos, nil
}
