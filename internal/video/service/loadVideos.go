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

// Load the videos
func (s *LoadVideosService) LoadVideos(req *video.LoadVideosRequest) ([]*base.Video, int64, error) {
	var lastTime int64
	var myId int64

	// If no lastTime in the argument, default is -1
	lastTime = -1
	if req.LastTime != nil {
		lastTime = *req.LastTime
	}
	myId = req.MyId

	// Fetch 30 videos based on the lastTime and create time order
	videoDBList, err := db.GetVideoListInOrder(lastTime)
	if err != nil {
		return nil, -1, err
	}

	// Initialize the return video list
	videoList := make([]*base.Video, len(videoDBList))

	// For each video record, fetch the author info and favorite info by rpc call
	for index, v := range videoDBList {
		// Fetch author info
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

		// Fetch favorite info, UserId == -1 means the user hasn't login, default favorite status is false
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

	// Extract the new lastTime to return
	if len(videoDBList) != 0 {
		lastTime = videoDBList[len(videoDBList)-1].CreateTime
	}

	return videoList, lastTime, nil
}
