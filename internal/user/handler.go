package main

import (
	"context"
	"github.com/CharmingCharm/DouSheng/idl/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	// TODO: Your code here...
	return
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
