package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/CharmingCharm/DouSheng/pkg/constants"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	Password      string `json:"password"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// CreateUser create user info
func CreateUser(ctx context.Context, name string, password string) (user_id int64, err error) {
	// TODO
	user := User{Name: name, Password: password, FollowCount: 0, FollowerCount: 0, IsFollow: false}
	res := DB.WithContext(ctx).Create(&user)
	if res.Error != nil {
		fmt.Println(res.Error.Error())
		return -1, res.Error
	}
	return int64(user.ID), nil
}

func GetUserByUsername(ctx context.Context, username string) (*User, error) {
	// TODO
	user := User{}
	fmt.Println(user)
	res := DB.Where(&User{Name: username}).First(&user)
	fmt.Println(user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		fmt.Println(res.Error.Error())
		return nil, res.Error
	}
	return &user, nil
}

// CreateUser create user info
func GetUserById(ctx context.Context, user_id int64) (*User, error) {
	// TODO
	return nil, nil
}

func GetUserListByIds(ctx context.Context, user_id []int64) ([]*User, error) {
	//TODO
	return nil, nil
}

func UserFollowCountPlus(ctx context.Context, user_id int64) error {
	// TODO
	return nil
}

func UserFollowCountSubtract(ctx context.Context, user_id int64) error {
	// TODO
	return nil
}

func UserFollowerCountPlus(ctx context.Context, user_id int64) error {
	// TODO
	return nil
}

func UserFollowerCountSubtract(ctx context.Context, user_id int64) error {
	// TODO
	return nil
}
