package service

import (
	"context"

	"github.com/CharmingCharm/DouSheng/pkg/status"

	"github.com/CharmingCharm/DouSheng/internal/user/db"
	"github.com/CharmingCharm/DouSheng/internal/user/rpc"
	"github.com/CharmingCharm/DouSheng/kitex_gen/action"
	"github.com/CharmingCharm/DouSheng/kitex_gen/base"
	"github.com/CharmingCharm/DouSheng/kitex_gen/user"
)

type GetUserInfoService struct {
	ctx context.Context
}

func NewGetUserInfoService(ctx context.Context) *GetUserInfoService {
	return &GetUserInfoService{ctx: ctx}
}

// Get user's info by its user id
func (s *GetUserInfoService) GetUserInfo(req *user.GetUserInfoRequest) (*base.User, error) {
	// Fetch the user info by its user id
	user_data, err := db.GetUserById(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	if user_data == nil {
		return nil, status.UserNotExistErr
	}

	// If current user is not logged in, the default relationship is not-follow
	flag := false
	if req.MyId != -1 {
		// Else, call rpc in action service to check relationship
		relationInfo, err := rpc.CheckRelation(s.ctx, &action.CheckRelationRequest{
			MyId:   req.MyId,
			UserId: user_data.Id,
		})
		if err != nil {
			return nil, err
		}
		if relationInfo.BaseResp.StatusCode != status.SuccessCode {
			return nil, status.NewStatus(relationInfo.BaseResp.StatusCode, relationInfo.BaseResp.StatusMessage)
		}
		flag = *relationInfo.IsFollow
	}

	// Return user
	user := base.User{
		Id:            user_data.Id,
		Name:          user_data.Name,
		FollowCount:   user_data.FollowCount,
		FollowerCount: user_data.FollowerCount,
		IsFollow:      flag,
	}
	return &user, nil
}
