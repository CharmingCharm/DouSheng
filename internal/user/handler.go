package main

import (
	"context"
	"fmt"

	"github.com/CharmingCharm/DouSheng/internal/user/pack"
	"github.com/CharmingCharm/DouSheng/internal/user/service"
	"github.com/CharmingCharm/DouSheng/kitex_gen/user"
	"github.com/CharmingCharm/DouSheng/pkg/status"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.CreateUserResponse)

	fmt.Println(req)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		fmt.Println("user/handler: 1")
		resp.BaseResp = pack.BuildBaseResp(status.ParamErr)
		return resp, nil
	}

	userId, err := service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		fmt.Println("user/handler: 2")
		resp.BaseResp = pack.BuildBaseResp(status.ConvertErrorToStatus(err))
		return resp, nil
	}
	fmt.Println("user/handler: 3")
	resp.UserId = userId
	resp.BaseResp = pack.BuildBaseResp(status.Success)
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserRequest) (resp *user.CheckUserResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (resp *user.GetUserInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateUserFollow implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUserFollow(ctx context.Context, req *user.UpdateUserFollowRequest) (resp *user.UpdateUserFollowResponse, err error) {
	// TODO: Your code here...
	return
}
