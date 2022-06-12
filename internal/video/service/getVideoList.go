package service

import (
	"context"

	"github.com/CharmingCharm/DouSheng/internal/video/db"
	"github.com/CharmingCharm/DouSheng/internal/video/rpc"
	"github.com/CharmingCharm/DouSheng/kitex_gen/action"
	"github.com/CharmingCharm/DouSheng/kitex_gen/base"
	"github.com/CharmingCharm/DouSheng/kitex_gen/user"
	"github.com/CharmingCharm/DouSheng/kitex_gen/video"
	"github.com/CharmingCharm/DouSheng/pkg/status"
)

type GetVideoListService struct {
	ctx context.Context
}

// NewGetVideoListService new GetVideoListService
func NewGetVideoListService(ctx context.Context) *GetVideoListService {
	return &GetVideoListService{ctx: ctx}
}

// Get the video list
func (s *GetVideoListService) GetVideoList(req *video.GetVideoListRequest) ([]*base.Video, error) {
	// Get the video list from video table based on video ids
	videoDBList, err := db.GetVideoListByIds(req.VideoIds)
	if err != nil {
		return nil, err
	}

	// Initialize return video list
	videoList := make([]*base.Video, len(videoDBList))

	// For each video record, fetch the author info and favorite info by rpc call
	for index, v := range videoDBList {
		// Fetch author info
		userInfo, err := rpc.GetUserInfo(s.ctx, &user.GetUserInfoRequest{
			UserId: v.AuthId,
			MyId:   req.UserId,
		})
		if err != nil {
			return nil, err
		}
		if userInfo.BaseResp.StatusCode != status.SuccessCode {
			return nil, status.NewStatus(userInfo.BaseResp.StatusCode, userInfo.BaseResp.StatusMessage)
		}

		// Fetch favorite info, UserId == -1 means the user hasn't login, default favorite status is false
		flag := false
		if req.UserId != -1 {
			favoriteInfo, err := rpc.CheckFavorite(s.ctx, &action.CheckFavoriteRequest{
				MyId:    req.UserId,
				VideoId: v.Id,
			})
			if err != nil {
				return nil, err
			}
			if favoriteInfo.BaseResp.StatusCode != status.SuccessCode {
				return nil, status.NewStatus(favoriteInfo.BaseResp.StatusCode, favoriteInfo.BaseResp.StatusMessage)
			}
			flag = *favoriteInfo.IsFavorite
		}

		videoList[index] = &base.Video{
			Id:            v.Id,
			Author:        userInfo.User,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    flag,
			Title:         v.Title,
		}
	}

	return videoList, nil
}
