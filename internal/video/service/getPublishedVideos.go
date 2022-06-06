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

	// UserId int64  `thrift:"user_id,1,required" json:"user_id"`
	// MyId   *int64 `thrift:"my_id,2" json:"my_id,omitempty"`

	videoDBList, err := db.GetVideoListByAuthorId(req.UserId)
	if err != nil {
		return nil, err
	}
	var myId int64
	myId = req.MyId

	videoList := make([]*base.Video, len(videoDBList))

	for index, v := range videoDBList {
		userInfo, err := rpc.GetUserInfo(s.ctx, &user.GetUserInfoRequest{
			UserId: v.AuthId,
			MyId:   myId,
		})
		if err != nil {
			return nil, err
		}
		if userInfo.BaseResp.StatusCode != status.SuccessCode {
			return nil, status.NewStatus(userInfo.BaseResp.StatusCode, userInfo.BaseResp.StatusMessage)
		}

		flag := false
		if myId != -1 {
			favoriteInfo, err := rpc.CheckFavorite(s.ctx, &action.CheckFavoriteRequest{
				MyId:    myId,
				VideoId: v.Id,
			})
			if err != nil {
				return nil, err
			}
			if favoriteInfo.BaseResp.StatusCode != status.SuccessCode {
				return nil, status.NewStatus(favoriteInfo.BaseResp.StatusCode, favoriteInfo.BaseResp.StatusMessage)
			}
			flag = favoriteInfo.IsFavorite
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
