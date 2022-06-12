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

type CreateUserService struct {
	ctx context.Context
}

func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser create user info.
func (s *CreateUserService) CreateUser(req *user.CreateUserRequest) (int64, error) {
	// Check if the user is already exist
	user, err := db.GetUserByUsername(s.ctx, req.Username)
	if err != nil {
		return -1, err
	}
	if user != nil {
		return -1, status.UserAlreadyExistErr
	}

	// Generate encrypted password and create user record
	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return -1, err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	var userId int64
	userId, err = db.CreateUser(s.ctx, req.Username, password)
	if err != nil {
		return -1, err
	}

	// Return id
	return userId, nil
}
