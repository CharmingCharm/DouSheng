package service

import (
	"context"

	// "github.com/CharmingCharm/DouSheng/internal/action/rpc"

	"github.com/CharmingCharm/DouSheng/internal/action/db"
	"github.com/CharmingCharm/DouSheng/kitex_gen/action"
)

type CheckFavoriteService struct {
	ctx context.Context
}

// NewCheckFavoriteService new CheckFavoriteService
func NewCheckFavoriteService(ctx context.Context) *CheckFavoriteService {
	return &CheckFavoriteService{ctx: ctx}
}

// CreateUser create user info.
func (s *CheckFavoriteService) CheckFavorite(req *action.CheckFavoriteRequest) (bool, error) {
	flag, err := db.CheckFavoriteRecord(s.ctx, req.MyId, req.VideoId)
	if err != nil {
		return false, err
	}
	return flag, nil
}
