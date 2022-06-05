package db

import (
	"github.com/CharmingCharm/DouSheng/pkg/constants"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Id            int64  `json:"id"`
	AuthId        int64  `json:"auth_id"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	CreateTime    int64  `json:"create_at" gorm:"autoCreateTime"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
}

func (v *Video) TableName() string {
	return constants.VideoTableName
}

func GetVideoListByIds(vIds []int64) ([]*Video, error) {
	// TODO
	return nil, nil
}

func GetVideoListByAuthorId(authorId int64) ([]*Video, error) {
	// TODO
	return nil, nil
}

func GetVideoListInOrder(lastTime int64) ([]*Video, error) {
	// TODO
	return nil, nil
}

func CreateVideo(authorId int64, playUrl string, coverUrl string) error {
	// TODO
	return nil
}

func VideoCommentCountAdd(vId int64) error {
	// TODO
	return nil
}

func VideoCommentCountSubtract(vId int64) error {
	// TODO
	return nil
}

func VideoFavoriteCountAdd(vId int64) error {
	// TODO
	return nil
}

func VideoFavoriteCountSubtract(vId int64) error {
	// TODO
	return nil
}
