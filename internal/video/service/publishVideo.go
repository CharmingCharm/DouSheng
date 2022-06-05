package service

import (
	"context"

	"github.com/CharmingCharm/DouSheng/kitex_gen/video"
)

type PublishVideoService struct {
	ctx context.Context
}

// NewPublishVideoService new PublishVideoService
func NewPublishVideoService(ctx context.Context) *PublishVideoService {
	return &PublishVideoService{ctx: ctx}
}

// CreateUser create user info.
func (s *PublishVideoService) PublishVideo(req *video.PublishVideoRequest) error {
	// TODO
	// db.CreateVideo
	return nil
}
