package service

import (
	"context"

	"github.com/CharmingCharm/DouSheng/pkg/constants"

	"github.com/CharmingCharm/DouSheng/internal/video/db"

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
	err := db.CreateVideo(s.ctx, req.MyId, req.DataUrl, constants.DefaultCoverUrl, req.Title)
	if err != nil {
		return err
	}

	return nil
}
