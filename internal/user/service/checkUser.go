package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/CharmingCharm/DouSheng/internal/user/db"
	"github.com/CharmingCharm/DouSheng/kitex_gen/user"
	"github.com/CharmingCharm/DouSheng/pkg/status"
)

type CheckUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{ctx: ctx}
}

// CreateUser create user info.
func (s *CheckUserService) CheckUser(req *user.CheckUserRequest) (int64, error) {
	user, err := db.GetUserByUsername(s.ctx, req.Username)
	if err != nil {
		return -1, err
	}
	if user == nil {
		return -1, status.LoginErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return -1, err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))

	if user.Password != password {
		return -1, status.LoginErr
	}
	return user.Id, nil
}
