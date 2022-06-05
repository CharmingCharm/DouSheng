package db

import (
	"context"
	"errors"

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
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// CreateUser create user info
func CreateUser(ctx context.Context, name string, password string) (user_id int64, err error) {
	// TODO
	user := User{Name: name, Password: password, FollowCount: 0, FollowerCount: 0}
	res := DB.WithContext(ctx).Create(&user)
	if res.Error != nil {
		return -1, res.Error
	}
	return int64(user.ID), nil
}

func GetUserByUsername(ctx context.Context, username string) (*User, error) {
	// TODO
	user := User{}
	res := DB.Where(&User{Name: username}).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return &user, nil
}

func GetUserById(ctx context.Context, user_id int64) (*User, error) {
	// TODO
	user := User{}
	res := DB.Where(&User{Id: user_id}).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return &user, nil
}

func GetUserListByIds(ctx context.Context, user_id []int64) ([]*User, error) {
	//TODO
	users := make([]*User, 0)
	res := DB.Find(&users, user_id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return users, nil
}

func UserFollowCountPlus(ctx context.Context, user_id int64) error {
	// TODO
	res := DB.Model(&User{}).Where("Id = ?", user_id).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1))
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func UserFollowCountSubtract(ctx context.Context, user_id int64) error {
	// TODO
	res := DB.Model(&User{}).Where("Id = ?", user_id).UpdateColumn("follow_count", gorm.Expr("follow_count - ?", 1))
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func UserFollowerCountPlus(ctx context.Context, user_id int64) error {
	// TODO
	res := DB.Model(&User{}).Where("Id = ?", user_id).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1))
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func UserFollowerCountSubtract(ctx context.Context, user_id int64) error {
	// TODO
	res := DB.Model(&User{}).Where("Id = ?", user_id).UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1))
	if res.Error != nil {
		return res.Error
	}
	return nil
}
