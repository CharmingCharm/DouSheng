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

func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{ctx: ctx}
}

// Check user's username and password
func (s *CheckUserService) CheckUser(req *user.CheckUserRequest) (int64, error) {
	// Get the user's info by username
	user, err := db.GetUserByUsername(s.ctx, req.Username)
	if err != nil {
		return -1, err
	}
	if user == nil {
		return -1, status.LoginErr
	}

	// Process the original password
	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return -1, err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))

	// Check password
	if user.Password != password {
		return -1, status.LoginErr
	}
	return user.Id, nil
}
