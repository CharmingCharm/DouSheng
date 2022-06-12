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
	// Based on the action type, do the add or delete comment count action
	if req.ActionType == 1 { // Add command:  add the count of comment of the video with 1
		err := db.VideoCommentCountAdd(req.VideoId)
		if err != nil {
			return err
		}
	} else if req.ActionType == 2 { // Delete command:  sutract the count of comment of the video with 1
		err := db.VideoCommentCountSubtract(req.VideoId)
		if err != nil {
			return err
		}
	} else {
		return status.ParamErr
	}
	return nil
}
