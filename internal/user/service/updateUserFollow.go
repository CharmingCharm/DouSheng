package service

import (
	"context"

	"github.com/CharmingCharm/DouSheng/internal/user/db"
	"github.com/CharmingCharm/DouSheng/kitex_gen/user"
	"github.com/CharmingCharm/DouSheng/pkg/status"
)

type UpdateUserFollowService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewUpdateUserFollowService(ctx context.Context) *UpdateUserFollowService {
	return &UpdateUserFollowService{ctx: ctx}
}

// CreateUser create user info.
func (s *UpdateUserFollowService) UpdateUserFollow(req *user.UpdateUserFollowRequest) error {

	fromUser, err1 := db.GetUserById(s.ctx, req.UserId)
	toUser, err2 := db.GetUserById(s.ctx, req.ToUserId)
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	if fromUser == nil || toUser == nil {
		return status.UserNotExistErr
	}

	var err error
	if req.ActionType == 1 {
		err = db.UserFollowCountPlus(s.ctx, req.UserId)
		if err != nil {
			return err
		}
		err = db.UserFollowerCountPlus(s.ctx, req.ToUserId)
		if err != nil {
			return err
		}
	} else if req.ActionType == 2 {
		db.UserFollowCountSubtract(s.ctx, req.UserId)
		if err != nil {
			return err
		}
		db.UserFollowerCountSubtract(s.ctx, req.ToUserId)
		if err != nil {
			return err
		}
	} else {
		return status.ActionTypeErr
	}
	return nil
}
