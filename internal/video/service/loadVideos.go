package service

import (
	"context"

	"github.com/CharmingCharm/DouSheng/internal/video/rpc"
	"github.com/CharmingCharm/DouSheng/pkg/status"

	"github.com/CharmingCharm/DouSheng/internal/video/db"

	"github.com/CharmingCharm/DouSheng/kitex_gen/action"
	"github.com/CharmingCharm/DouSheng/kitex_gen/base"
	"github.com/CharmingCharm/DouSheng/kitex_gen/user"
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

	var lastTime int64
	var myId int64

	lastTime = -1
	if req.LastTime != nil {
		lastTime = *req.LastTime
	}
	myId = req.MyId

	videoDBList, err := db.GetVideoListInOrder(lastTime)
	videoList := make([]*base.Video, len(videoDBList))

	if err != nil {
		return nil, -1, err
	}

	for index, v := range videoDBList {
		userInfo, err := rpc.GetUserInfo(s.ctx, &user.GetUserInfoRequest{
			UserId: v.AuthId,
			MyId:   myId,
		})
		if err != nil {
			return nil, -1, err
		}
		if userInfo.BaseResp.StatusCode != status.SuccessCode {
			return nil, -1, status.NewStatus(userInfo.BaseResp.StatusCode, userInfo.BaseResp.StatusMessage)
		}

		flag := false
		if myId != -1 {
			favoriteInfo, err := rpc.CheckFavorite(s.ctx, &action.CheckFavoriteRequest{
				MyId:    myId,
				VideoId: v.Id,
			})
			if err != nil {
				return nil, -1, err
			}
			if favoriteInfo.BaseResp.StatusCode != status.SuccessCode {
				return nil, -1, status.NewStatus(favoriteInfo.BaseResp.StatusCode, favoriteInfo.BaseResp.StatusMessage)
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

	if len(videoDBList) != 0 {
		lastTime = videoDBList[len(videoDBList)-1].CreateTime
	}

	return videoList, lastTime, nil
}
