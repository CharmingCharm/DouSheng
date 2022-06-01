// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser create user info.
func (s *CreateUserService) CreateUser(req *user.CreateUserRequest) (int64, error) {
	user, err := db.GetUserByUsername(s.ctx, req.Username)
	if err != nil {
		return -1, err
	}
	if user != nil {
		return -1, status.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return -1, err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	var userId int64
	userId, err = db.CreateUser(s.ctx, req.Username, password)
	if err != nil {
		return -1, nil
	}
	return userId, err
}
