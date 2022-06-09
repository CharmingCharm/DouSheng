package db

import (
	"context"

	"github.com/CharmingCharm/DouSheng/pkg/status"

	"github.com/CharmingCharm/DouSheng/pkg/constants"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Id         int64  `json:"id"`
	UserId     int64  `json:"user_id"`
	VideoId    int64  `json:"video_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time" gorm:"autoCreateTime"`
}

func (c *Comment) TableName() string {
	return constants.CommentTableName
}

func GetCommentListByVideoId(ctx context.Context, videoId int64) ([]*Comment, error) {
	// TODO
	commentList := make([]*Comment, 0)
	res := DB.Where("video_id = ?", videoId).Order("create_time desc").Find(&commentList)
	if res.Error != nil {
		return nil, res.Error
	}
	return commentList, nil
}

func CreateCommentRecord(ctx context.Context, userId int64, videoId int64, commentText string) (Comment, error) {
	// TODO
	comment := Comment{
		UserId:  userId,
		VideoId: videoId,
		Content: commentText,
	}
	res := DB.WithContext(ctx).Create(&comment)
	if res.Error != nil {
		return comment, res.Error
	}
	return comment, nil
}

func DeleteCommentRecord(ctx context.Context, userId int64, videoId int64, commentId int64) error {
	// TODO
	comment := Comment{}
	searchRes := DB.Model(&Comment{}).Where(&Comment{Id: commentId})
	searchRes.First(&comment)

	if comment.UserId != userId || comment.VideoId != videoId {
		return status.ParamControdictErr
	}

	res := searchRes.Delete(&Comment{})

	if res.Error != nil {
		return res.Error
	}
	return nil
}
