package db

import (
	"context"
	"fmt"
	"time"

	"github.com/CharmingCharm/DouSheng/pkg/constants"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Id            int64  `json:"id"`
	AuthId        int64  `json:"auth_id"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	CreateTime    int64  `json:"create_time" gorm:"autoCreateTime"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	Title         string `json:"title"`
}

func (v *Video) TableName() string {
	return constants.VideoTableName
}

func GetVideoListByIds(vIds []int64) ([]*Video, error) {
	// TODO
	videoList := make([]*Video, 0)
	if len(vIds) == 0 {
		return videoList, nil
	}

	res := DB.Model(&Video{}).Where("id IN ?", vIds).Order("create_time desc").Find(&videoList)
	if res.Error != nil {
		return nil, res.Error
	}
	return videoList, nil
}

func GetVideoListByAuthorId(authorId int64) ([]*Video, error) {
	// TODO
	videoList := make([]*Video, 0)
	res := DB.Model(&Video{}).Where("auth_id = ?", authorId).Order("create_time desc").Find(&videoList)
	if res.Error != nil {
		return nil, res.Error
	}
	return videoList, nil
}

func GetVideoListInOrder(lastTime int64) ([]*Video, error) {
	// TODO
	videoList := make([]*Video, 0)
	if lastTime == -1 {
		res := DB.Model(&Video{}).Order("create_time desc").Limit(30).Find(&videoList)

		if res.Error != nil {
			return nil, res.Error
		}
		return videoList, nil
	}
	res := DB.Where("create_time < ?", lastTime).Order("create_time desc").Limit(30).Find(&videoList)
	if res.Error != nil {
		return nil, res.Error
	}
	return videoList, nil
}

func CreateVideo(ctx context.Context, authorId int64, playUrl string, coverUrl string, title string) error {
	// TODO
	video := Video{
		AuthId:        authorId,
		PlayUrl:       playUrl,
		CoverUrl:      coverUrl,
		CreateTime:    time.Now().Unix(),
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
	}
	res := DB.WithContext(ctx).Create(&video)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func VideoCommentCountAdd(vId int64) error {
	// TODO
	res := DB.Model(&Video{}).Where("Id = ?", vId).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1))
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func VideoCommentCountSubtract(vId int64) error {
	// TODO
	res := DB.Model(&Video{}).Where("Id = ?", vId).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1))
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func VideoFavoriteCountAdd(vId int64) error {
	// TODO
	fmt.Println(vId)
	res := DB.Model(&Video{}).Where("Id = ?", vId).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1))
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func VideoFavoriteCountSubtract(vId int64) error {
	// TODO
	res := DB.Model(&Video{}).Where("Id = ?", vId).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1))
	if res.Error != nil {
		return res.Error
	}
	return nil
}
