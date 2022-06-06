package db

import (
	"context"
	"errors"

	"github.com/CharmingCharm/DouSheng/pkg/constants"

	"gorm.io/gorm"
)

type Relation struct {
	gorm.Model
	Id       int64 `json:"id"`
	UserId   int64 `json:"user_id"`
	ToUserId int64 `json:"to_user_id"`
}

func (r *Relation) TableName() string {
	return constants.RelationTableName
}

func FindRelationRecord(ctx context.Context, userId int64, toUserId int64) (bool, error) {
	// TODO
	res := DB.Where(&Relation{UserId: userId, ToUserId: toUserId}).First(&Relation{})
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, res.Error
	}
	return true, nil
}

func GetFollowerList(ctx context.Context, toUserId int64) ([]int64, error) {
	// TODO
	followerList := make([]int64, 0)
	res := DB.Where(&Relation{ToUserId: toUserId}).Select("user_id").Find(&followerList)
	if res.Error != nil {
		return nil, res.Error
	}
	return followerList, nil
}

func GetFollowList(ctx context.Context, userId int64) ([]int64, error) {
	// TODO
	followList := make([]int64, 0)
	res := DB.Where(&Relation{UserId: userId}).Select("to_user_id").Find(&followList)
	if res.Error != nil {
		return nil, res.Error
	}
	return followList, nil
}

func CreateRelationshipRecord(ctx context.Context, userId int64, toUserId int64) error {
	// TODO
	relation := Relation{
		UserId:   userId,
		ToUserId: toUserId,
	}
	res := DB.WithContext(ctx).Create(&relation)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func DeleteRelationshipRecord(ctx context.Context, userId int64, toUserId int64) error {
	// TODO
	res := DB.Where(&Relation{UserId: userId, ToUserId: toUserId}).Delete(&Relation{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
