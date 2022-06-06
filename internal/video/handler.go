package main

import (
	"context"

	"github.com/CharmingCharm/DouSheng/internal/video/service"
	"github.com/CharmingCharm/DouSheng/kitex_gen/video"
	"github.com/CharmingCharm/DouSheng/pkg/response"
	"github.com/CharmingCharm/DouSheng/pkg/status"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// GetVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoList(ctx context.Context, req *video.GetVideoListRequest) (resp *video.GetVideoListResponse, err error) {
	// TODO: Your code here...
	resp = new(video.GetVideoListResponse)

	if len(req.VideoIds) == 0 {
		resp.BaseResp = response.BuildBaseResp(status.ParamErr)
		return resp, nil
	}

	videoList, err := service.NewGetVideoListService(ctx).GetVideoList(req)
	if err != nil {
		resp.BaseResp = response.BuildBaseResp(status.ConvertErrorToStatus(err))
		return resp, nil
	}
	resp.BaseResp = response.BuildBaseResp(status.Success)
	resp.Videos = videoList
	return resp, nil
}

// LoadVideos implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) LoadVideos(ctx context.Context, req *video.LoadVideosRequest) (resp *video.LoadVideosResponse, err error) {
	// TODO: Your code here...
	resp = new(video.LoadVideosResponse)

	videoList, nextTime, err := service.NewLoadVideosService(ctx).LoadVideos(req)
	if err != nil {
		resp.BaseResp = response.BuildBaseResp(status.ConvertErrorToStatus(err))
		return resp, nil
	}
	resp.BaseResp = response.BuildBaseResp(status.Success)
	resp.VideoList = videoList
	if nextTime != -1 {
		resp.NextTime = &nextTime
		// fmt.Println("lala")
	}
	return resp, nil
}

// PublishVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishVideo(ctx context.Context, req *video.PublishVideoRequest) (resp *video.PublishVideoResponse, err error) {
	// TODO: Your code here...
	resp = new(video.PublishVideoResponse)

	if len(req.DataUrl) == 0 || len(req.Title) == 0 {
		resp.BaseResp = response.BuildBaseResp(status.ParamErr)
		return resp, nil
	}

	err = service.NewPublishVideoService(ctx).PublishVideo(req)
	if err != nil {
		resp.BaseResp = response.BuildBaseResp(status.ConvertErrorToStatus(err))
		return resp, nil
	}
	resp.BaseResp = response.BuildBaseResp(status.Success)
	return resp, nil
}

// GetPublishedVideos implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishedVideos(ctx context.Context, req *video.GetPublishedVideosRequest) (resp *video.GetPublishedVideosResponse, err error) {
	// TODO: Your code here...
	resp = new(video.GetPublishedVideosResponse)

	if req.UserId <= 0 {
		resp.BaseResp = response.BuildBaseResp(status.ParamErr)
		return resp, nil
	}

	videoList, err := service.NewGetPublishedVideosService(ctx).GetPublishedVideos(req)
	if err != nil {
		resp.BaseResp = response.BuildBaseResp(status.ConvertErrorToStatus(err))
		return resp, nil
	}
	resp.BaseResp = response.BuildBaseResp(status.Success)
	resp.VideoList = videoList
	return resp, nil
}

// UpdateFavoriteCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) UpdateFavoriteCount(ctx context.Context, req *video.UpdateFavoriteCountRequest) (resp *video.UpdateFavoriteCountResponse, err error) {
	// TODO: Your code here...
	resp = new(video.UpdateFavoriteCountResponse)

	if req.VideoId <= 0 || (req.ActionType != 1 && req.ActionType != 2) {
		resp.BaseResp = response.BuildBaseResp(status.ParamErr)
		return resp, nil
	}

	err = service.NewUpdateFavoriteCountService(ctx).UpdateFavoriteCount(req)
	if err != nil {
		resp.BaseResp = response.BuildBaseResp(status.ConvertErrorToStatus(err))
		return resp, nil
	}
	resp.BaseResp = response.BuildBaseResp(status.Success)
	return resp, nil
}

// UpdateCommentCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) UpdateCommentCount(ctx context.Context, req *video.UpdateCommentCountRequest) (resp *video.UpdateCommentCountResponse, err error) {
	// TODO: Your code here...
	resp = new(video.UpdateCommentCountResponse)

	if req.VideoId <= 0 || (req.ActionType != 1 && req.ActionType != 2) {
		resp.BaseResp = response.BuildBaseResp(status.ParamErr)
		return resp, nil
	}

	err = service.NewUpdateCommentCountService(ctx).UpdateCommentCount(req)
	if err != nil {
		resp.BaseResp = response.BuildBaseResp(status.ConvertErrorToStatus(err))
		return resp, nil
	}
	resp.BaseResp = response.BuildBaseResp(status.Success)
	return resp, nil
}
