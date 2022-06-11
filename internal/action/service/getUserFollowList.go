package service

import (
	"context"

	// "github.com/CharmingCharm/DouSheng/internal/action/rpc"

	"github.com/CharmingCharm/DouSheng/internal/action/db"
	"github.com/CharmingCharm/DouSheng/internal/action/rpc"
	"github.com/CharmingCharm/DouSheng/kitex_gen/action"
	"github.com/CharmingCharm/DouSheng/kitex_gen/base"
	"github.com/CharmingCharm/DouSheng/kitex_gen/user"
	"github.com/CharmingCharm/DouSheng/pkg/status"
)

type GetUserFollowListService struct {
	ctx context.Context
}

// NewGetUserFollowListService new GetUserFollowListService
func NewGetUserFollowListService(ctx context.Context) *GetUserFollowListService {
	return &GetUserFollowListService{ctx: ctx}
}

// Get the follow list
func (s *GetUserFollowListService) GetUserFollowList(req *action.GetUserFollowListRequest) ([]*base.User, error) {
	followDBList, err := db.GetFollowList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	userList := make([]*base.User, len(followDBList))

	for index, f := range followDBList {
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
