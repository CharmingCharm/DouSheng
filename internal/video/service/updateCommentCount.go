package service

import (
	"context"

	"github.com/CharmingCharm/DouSheng/internal/video/db"
	"github.com/CharmingCharm/DouSheng/pkg/status"

	"github.com/CharmingCharm/DouSheng/kitex_gen/video"
)

type UpdateCommentCountService struct {
	ctx context.Context
}

// NewUpdateCommentCountService new UpdateCommentCountService
func NewUpdateCommentCountService(ctx context.Context) *UpdateCommentCountService {
	return &UpdateCommentCountService{ctx: ctx}
}

// Update the number of comments
func (s *UpdateCommentCountService) UpdateCommentCount(req *video.UpdateCommentCountRequest) error {
	if req.ActionType == 1 {
		err := db.VideoCommentCountAdd(req.VideoId)
		if err != nil {
			return err
		}
	} else if req.ActionType == 2 {
		err := db.VideoCommentCountSubtract(req.VideoId)
		if err != nil {
			return err
		}
	} else {
		return status.ParamErr
	}
	return nil
}
