package service

import (
	"context"

	"github.com/CharmingCharm/DouSheng/internal/action/db"
	"github.com/CharmingCharm/DouSheng/internal/action/rpc"
	"github.com/CharmingCharm/DouSheng/pkg/status"

	"github.com/CharmingCharm/DouSheng/kitex_gen/action"
	"github.com/CharmingCharm/DouSheng/kitex_gen/video"
)

type UpdateCommentService struct {
	ctx context.Context
}

// NewUpdateCommentService new UpdateCommentService d
func NewUpdateCommentService(ctx context.Context) *UpdateCommentService {
	return &UpdateCommentService{ctx: ctx}
}

// CreateUser create user info.
func (s *UpdateCommentService) UpdateComment(req *action.UpdateCommentRequest) error {

	// type UpdateCommentRequest struct {
	// 	UserId      int64   `thrift:"user_id,1,required" json:"user_id"`
	// 	VideoId     int64   `thrift:"video_id,2,required" json:"video_id"`
	// 	ActionType  int32   `thrift:"action_type,3,required" json:"action_type"`
	// 	CommentText *string `thrift:"comment_text,4" json:"comment_text,omitempty"`
	// 	CommentId   *int64  `thrift:"comment_id,5" json:"comment_id,omitempty"`
	// }

	// db.CreateCommentRecord()
	// db.DeleteCommentRecord()

	// rpc.video.UpdateCommentCount()

	if req.ActionType == 1 {
		if req.CommentText == nil {
			return status.ParamErr
		}
		err := db.CreateCommentRecord(s.ctx, req.UserId, req.VideoId, *req.CommentText)
		if err != nil {
			return err
		}
	} else if req.ActionType == 2 {
		if req.CommentId == nil {
			return status.ParamErr
		}
		err := db.DeleteCommentRecord(s.ctx, req.UserId, req.VideoId, *req.CommentId)
		if err != nil {
			return err
		}
	}

	resp, err := rpc.UpdateCommentCount(s.ctx, &video.UpdateCommentCountRequest{
		VideoId:    req.VideoId,
		ActionType: req.ActionType,
	})
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != status.SuccessCode {
		return status.NewStatus(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}
