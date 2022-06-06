package service

import (
	"context"

	"github.com/CharmingCharm/DouSheng/internal/action/db"
	"github.com/CharmingCharm/DouSheng/internal/action/rpc"
	"github.com/CharmingCharm/DouSheng/pkg/status"

	"github.com/CharmingCharm/DouSheng/kitex_gen/action"
	"github.com/CharmingCharm/DouSheng/kitex_gen/base"
	"github.com/CharmingCharm/DouSheng/kitex_gen/user"
)

type GetUserFollowerListService struct {
	ctx context.Context
}

// NewGetUserFollowerListService new GetUserFollowerListService
func NewGetUserFollowerListService(ctx context.Context) *GetUserFollowerListService {
	return &GetUserFollowerListService{ctx: ctx}
}

// CreateUser create user info.
func (s *GetUserFollowerListService) GetUserFollowerList(req *action.GetUserFollowerListRequest) ([]*base.User, error) {

	// type GetUserFollowerListRequest struct {
	// 	UserId int64 `thrift:"user_id,1,required" json:"user_id"`
	// 	MyId   int64 `thrift:"my_id,2,required" json:"my_id"`
	// }

	followerDBList, err := db.GetFollowerList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	userList := make([]*base.User, len(followerDBList))

	for index, f := range followerDBList {
		resp, err := rpc.GetUserInfo(s.ctx, &user.GetUserInfoRequest{
			UserId: f,
			MyId:   req.MyId,
		})
		if err != nil {
			return nil, err
		}
		if resp.BaseResp.StatusCode != status.SuccessCode {
			return nil, status.NewStatus(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
		}

		userList[index] = resp.User
	}
	return userList, nil
}
