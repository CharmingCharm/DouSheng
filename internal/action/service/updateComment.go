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
	if req.ActionType == 1 {
		if req.CommentText == nil {
			return nil, status.ParamErr
		}

		commentDB, err := db.CreateCommentRecord(s.ctx, req.UserId, req.VideoId, *req.CommentText)
		if err != nil {
			return nil, err
		}

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

		comment = base.Comment{
			Id:         commentDB.Id,
			User:       userInfo.User,
			Content:    commentDB.Content,
			CreateDate: time.Unix(commentDB.CreateTime, 0).Format("2006-01-02 15:04"),
		}
	} else if req.ActionType == 2 {
		if req.CommentId == nil {
			return nil, status.ParamErr
		}
		err := db.DeleteCommentRecord(s.ctx, req.UserId, req.VideoId, *req.CommentId)
		if err != nil {
			return nil, err
		}
	}

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

	if req.ActionType == 2 {
		return nil, nil
	}
	return &comment, nil
}
