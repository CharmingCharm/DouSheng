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

// Get the list of fans
func (s *GetUserFollowerListService) GetUserFollowerList(req *action.GetUserFollowerListRequest) ([]*base.User, error) {
	// Get the follower's user id by the req.UserId from relation table
	followerDBList, err := db.GetFollowerList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	// Initialize the return user list
	userList := make([]*base.User, len(followerDBList))

	// For each fetch follower's id, do rpc call to fetch follower's info from user server
	for index, f := range followerDBList {
		// Rpc call for follower's info
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

	// Return the user list
	return userList, nil
}
