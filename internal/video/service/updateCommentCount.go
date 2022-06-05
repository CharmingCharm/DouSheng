package service

import (
	"context"

	"github.com/CharmingCharm/DouSheng/kitex_gen/video"
)

type UpdateCommentCountService struct {
	ctx context.Context
}

// NewUpdateCommentCountService new UpdateCommentCountService
func NewUpdateCommentCountService(ctx context.Context) *UpdateCommentCountService {
	return &UpdateCommentCountService{ctx: ctx}
}

// CreateUser create user info.
func (s *UpdateCommentCountService) UpdateCommentCount(req *video.UpdateCommentCountRequest) error {
	// db.VideoCommentCountAdd
	// db.VideoCommentCountSubtract

	// for v in videoList { rpc.user.GetUserInfo }
	return nil
}
