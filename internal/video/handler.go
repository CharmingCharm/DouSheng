package main

import (
	"context"
	"github.com/CharmingCharm/DouSheng/idl/kitex_gen/video"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// GetVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoList(ctx context.Context, req *video.GetVideoListRequest) (resp *video.GetVideoListResponse, err error) {
	// TODO: Your code here...
	return
}

// LoadVideos implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) LoadVideos(ctx context.Context, req *video.LoadVideosRequest) (resp *video.LoadVideosResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishVideo(ctx context.Context, req *video.PublishVideoRequest) (resp *video.PublishVideoResponse, err error) {
	// TODO: Your code here...
	return
}

// GetPublishedVideos implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishedVideos(ctx context.Context, req *video.GetPublishedVideosRequest) (resp *video.GetPublishedVideosResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateFavoriteCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) UpdateFavoriteCount(ctx context.Context, req *video.UpdateFavoriteCountRequest) (resp *video.UpdateFavoriteCountResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateCommentCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) UpdateCommentCount(ctx context.Context, req *video.UpdateCommentCountRequest) (resp *video.UpdateCommentCountResponse, err error) {
	// TODO: Your code here...
	return
}
