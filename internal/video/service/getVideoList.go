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

// CreateUser create user info.
func (s *GetVideoListService) GetVideoList(req *video.GetVideoListRequest) ([]*base.Video, error) {
	// videoList := db.GetVideoListByIds
	// VideoIds []int64 `thrift:"video_ids,1,required" json:"video_ids"`
	// UserId   int64   `thrift:"user_id,2,required" json:"user_id"`

	videoDBList, err := db.GetVideoListByIds(req.VideoIds)
	if err != nil {
		return nil, err
	}

	videoList := make([]*base.Video, len(videoDBList))

	for index, v := range videoDBList {
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
			IsFavorite:    flag, // TODO
			Title:         v.Title,
		}
	}

	// for v in videoList { rpc.user.GetUserInfo }
	return videoList, nil
}
