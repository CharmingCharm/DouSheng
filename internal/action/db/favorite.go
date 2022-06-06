package db

import (
	"context"
	"errors"

	"github.com/CharmingCharm/DouSheng/pkg/constants"

	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	Id      int64 `json:"id"`
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
}

func (f *Favorite) TableName() string {
	return constants.FavorateTableName
}

func CheckFavoriteRecord(cts context.Context, userId int64, videoId int64) (bool, error) {
	res := DB.Model(&Favorite{}).Where(&Favorite{UserId: userId, VideoId: videoId}).First(&Favorite{})
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, res.Error
	}
	return true, nil
}

func GetFavoriteVideoIdsByUserId(ctx context.Context, userId int64) ([]int64, error) {
	// TODO
	vIds := make([]int64, 0)
	res := DB.Model(&Favorite{}).Where(&Favorite{UserId: userId}).Select("video_id").Find(&vIds)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return vIds, nil
}

func CreateFavoriteRecord(ctx context.Context, userId int64, videoId int64) error {
	// TODO
	favorite := Favorite{
		UserId:  userId,
		VideoId: videoId,
	}
	res := DB.WithContext(ctx).Create(&favorite)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func DeleteFavoriteRecord(ctx context.Context, userId int64, videoId int64) error {
	// TODO
	res := DB.Model(&Favorite{}).Where("user_id = ? AND video_id = ?", userId, videoId).Delete(&Favorite{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
