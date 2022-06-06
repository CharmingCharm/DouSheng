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

// CreateUser create user info.
func (s *GetCommentListService) GetCommentList(req *action.GetCommentListRequest) ([]*base.Comment, error) {

	// type GetCommentListRequest struct {
	// 	MyId    int64 `thrift:"my_id,2,required" json:"my_id"`
	// 	VideoId int64 `thrift:"video_id,3,required" json:"video_id"`
	// }

	commentDBList, err := db.GetCommentListByVideoId(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	commentList := make([]*base.Comment, len(commentDBList))

	for index, c := range commentDBList {
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
			CreateDate: time.Unix(c.CreateTime, 0).Format("2006-01-02 15:04"),
		}
	}

	// for c := range commentList { rpc.user.GetUserInfo() }
	return commentList, nil
}
