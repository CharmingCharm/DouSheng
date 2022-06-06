package main

import (
	"context"

	"github.com/CharmingCharm/DouSheng/internal/user/service"
	"github.com/CharmingCharm/DouSheng/kitex_gen/user"
	"github.com/CharmingCharm/DouSheng/pkg/response"
	"github.com/CharmingCharm/DouSheng/pkg/status"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.CreateUserResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = response.BuildBaseResp(status.ParamErr)
		return resp, nil
	}

	userId, err := service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = response.BuildBaseResp(status.ConvertErrorToStatus(err))
		return resp, nil
	}
	resp.BaseResp = response.BuildBaseResp(status.Success)
	resp.UserId = &userId
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserRequest) (resp *user.CheckUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.CheckUserResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = response.BuildBaseResp(status.ParamErr)
		return resp, nil
	}

	userId, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = response.BuildBaseResp(status.ConvertErrorToStatus(err))
		return resp, nil
	}
	resp.BaseResp = response.BuildBaseResp(status.Success)
	resp.UserId = &userId
	return resp, nil
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (resp *user.GetUserInfoResponse, err error) {
	// TODO: Your code here...
	resp = new(user.GetUserInfoResponse)

	user, err := service.NewGetUserInfoService(ctx).GetUserInfo(req)
	if err != nil {
		resp.BaseResp = response.BuildBaseResp(status.ConvertErrorToStatus(err))
		return resp, nil
	}
	resp.BaseResp = response.BuildBaseResp(status.Success)
	resp.User = user
	return resp, nil
}

// UpdateUserFollow implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUserFollow(ctx context.Context, req *user.UpdateUserFollowRequest) (resp *user.UpdateUserFollowResponse, err error) {
	// TODO: Your code here...
	resp = new(user.UpdateUserFollowResponse)

	err = service.NewUpdateUserFollowService(ctx).UpdateUserFollow(req)
	if err != nil {
		resp.BaseResp = response.BuildBaseResp(status.ConvertErrorToStatus(err))
		return resp, nil
	}
	resp.BaseResp = response.BuildBaseResp(status.Success)
	return resp, nil
}
