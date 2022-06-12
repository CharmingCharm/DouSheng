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
)

type GetCommentListService struct {
	ctx context.Context
}

// NewGetCommentListService new GetCommentListService
func NewGetCommentListService(ctx context.Context) *GetCommentListService {
	return &GetCommentListService{ctx: ctx}
}

// Get the list of video comments
func (s *GetCommentListService) GetCommentList(req *action.GetCommentListRequest) ([]*base.Comment, error) {
	// Get the comment list of a video by the video id
	commentDBList, err := db.GetCommentListByVideoId(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}

	// Initialize the return comment list
	commentList := make([]*base.Comment, len(commentDBList))

	// For each comment record, fetch its author info by author id
	for index, c := range commentDBList {
		// Fetch comment's author by author id
		resp, err := rpc.GetUserInfo(s.ctx, &user.GetUserInfoRequest{
			UserId: c.UserId,
			MyId:   req.MyId,
		})
		if err != nil {
			return nil, err
		}
		if resp.BaseResp.StatusCode != status.SuccessCode {
			return nil, status.NewStatus(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
		}

		commentList[index] = &base.Comment{
			Id:         c.Id,
			User:       resp.User,
			Content:    c.Content,
			CreateDate: time.Unix(c.CreateTime, 0).Format("2006-01-02 15:04"), // Transfer the create time in data base to readable type
		}
	}

	return commentList, nil
}
