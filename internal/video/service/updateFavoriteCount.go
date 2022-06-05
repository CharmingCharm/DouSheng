package service

import (
	"context"

	"github.com/CharmingCharm/DouSheng/kitex_gen/video"
)

type UpdateFavoriteCountService struct {
	ctx context.Context
}

// NewUpdateFavoriteCountService new UpdateFavoriteCountService
func NewUpdateFavoriteCountService(ctx context.Context) *UpdateFavoriteCountService {
	return &UpdateFavoriteCountService{ctx: ctx}
}

// CreateUser create user info.
func (s *UpdateFavoriteCountService) UpdateFavoriteCount(req *video.UpdateFavoriteCountRequest) error {
	// db.VideoFavoriteCountAdd
	// db.VideoFavoriteCountSubtract

	// for v in videoList { rpc.user.GetUserInfo }
	return nil
}
