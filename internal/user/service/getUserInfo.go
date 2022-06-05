package service

import (
	"context"

	"github.com/CharmingCharm/DouSheng/pkg/status"

	"github.com/CharmingCharm/DouSheng/internal/user/db"
	"github.com/CharmingCharm/DouSheng/kitex_gen/base"
	"github.com/CharmingCharm/DouSheng/kitex_gen/user"
)

type GetUserInfoService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewGetUserInfoService(ctx context.Context) *GetUserInfoService {
	return &GetUserInfoService{ctx: ctx}
}

// CreateUser create user info.
func (s *GetUserInfoService) GetUserInfo(req *user.GetUserInfoRequest) (*base.User, error) {
	user_data, err := db.GetUserById(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	if user_data == nil {
		return nil, status.UserNotExistErr
	}

	user := base.User{
		Id:            user_data.Id,
		Name:          user_data.Name,
		FollowCount:   user_data.FollowCount,
		FollowerCount: user_data.FollowerCount,
		IsFollow:      false,
	}
	// Call rpc.action.CheckRelation
	return &user, nil
}
