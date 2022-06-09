package service

import (
	"context"

	"github.com/CharmingCharm/DouSheng/internal/action/db"

	"github.com/CharmingCharm/DouSheng/internal/action/rpc"

	"github.com/CharmingCharm/DouSheng/kitex_gen/action"
	"github.com/CharmingCharm/DouSheng/kitex_gen/video"
	"github.com/CharmingCharm/DouSheng/pkg/status"
)

type UpdateFavoriteService struct {
	ctx context.Context
}

// NewUpdateFavoriteService new UpdateFavoriteService
func NewUpdateFavoriteService(ctx context.Context) *UpdateFavoriteService {
	return &UpdateFavoriteService{ctx: ctx}
}

// CreateUser create user info.
func (s *UpdateFavoriteService) UpdateFavorite(req *action.UpdateFavoriteRequest) error {
	flag, err := db.CheckFavoriteRecord(context.Background(), req.UserId, req.VideoId)
	if err != nil {
		return err
	}

	if req.ActionType == 1 {
		if flag {
			return nil
		}
		err := db.CreateFavoriteRecord(context.Background(), req.UserId, req.VideoId)
		if err != nil {
			return err
		}
	} else if req.ActionType == 2 {
		if !flag {
			return nil
		}
		err := db.DeleteFavoriteRecord(context.Background(), req.UserId, req.VideoId)
		if err != nil {
			return err
		}
	}

	resp, err := rpc.UpdateFavoriteCount(context.Background(), &video.UpdateFavoriteCountRequest{
		VideoId:    req.VideoId,
		ActionType: req.ActionType,
	})
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != status.SuccessCode {
		return status.NewStatus(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}
