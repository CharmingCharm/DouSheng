package service

import (
	"context"
	"time"

	"github.com/CharmingCharm/DouSheng/internal/action/db"
	"github.com/CharmingCharm/DouSheng/internal/action/rpc"
	"github.com/CharmingCharm/DouSheng/pkg/status"

	"github.com/CharmingCharm/DouSheng/kitex_gen/action"
	"github.com/CharmingCharm/DouSheng/kitex_gen/base"
	"github.com/CharmingCharm/DouSheng/kitex_gen/user"
	"github.com/CharmingCharm/DouSheng/kitex_gen/video"
)

type UpdateCommentService struct {
	ctx context.Context
}

// NewUpdateCommentService new UpdateCommentService d
func NewUpdateCommentService(ctx context.Context) *UpdateCommentService {
	return &UpdateCommentService{ctx: ctx}
}

// Update the comments
func (s *UpdateCommentService) UpdateComment(req *action.UpdateCommentRequest) (*base.Comment, error) {
	var comment base.Comment
	// Check action type to create a comment record or to delete a comment record
	if req.ActionType == 1 { // Create comment
		// Check CommentText argument
		if req.CommentText == nil {
			return nil, status.ParamErr
		}
		// Create comment record in comment table
		commentDB, err := db.CreateCommentRecord(s.ctx, req.UserId, req.VideoId, *req.CommentText)
		if err != nil {
			return nil, err
		}

		// Fetch author info by rpc call
		userInfo, err := rpc.GetUserInfo(s.ctx, &user.GetUserInfoRequest{
			UserId: commentDB.UserId,
			MyId:   req.UserId,
		})
		if err != nil {
			return nil, err
		}
		if userInfo.BaseResp.StatusCode != status.SuccessCode {
			return nil, status.NewStatus(userInfo.BaseResp.StatusCode, userInfo.BaseResp.StatusMessage)
		}

		// Attached the author's info
		comment = base.Comment{
			Id:         commentDB.Id,
			User:       userInfo.User,
			Content:    commentDB.Content,
			CreateDate: time.Unix(commentDB.CreateTime, 0).Format("2006-01-02 15:04"),
		}
	} else if req.ActionType == 2 { // Delete comment
		// Check CommentId
		if req.CommentId == nil {
			return nil, status.ParamErr
		}
		// Delete the corresponding omment
		err := db.DeleteCommentRecord(s.ctx, req.UserId, req.VideoId, *req.CommentId)
		if err != nil {
			return nil, err
		}
	}

	// Update comment count by rpc call to video service
	resp, err := rpc.UpdateCommentCount(s.ctx, &video.UpdateCommentCountRequest{
		VideoId:    req.VideoId,
		ActionType: req.ActionType,
	})
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != status.SuccessCode {
		return nil, status.NewStatus(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	// Do response based on action type
	if req.ActionType == 2 {
		return nil, nil
	}
	return &comment, nil
}
