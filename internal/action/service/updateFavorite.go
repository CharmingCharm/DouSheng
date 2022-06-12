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

// Update the "Favorite"
func (s *UpdateFavoriteService) UpdateFavorite(req *action.UpdateFavoriteRequest) error {
	// Check if the favorite record exists before update
	flag, err := db.CheckFavoriteRecord(s.ctx, req.UserId, req.VideoId)
	if err != nil {
		return err
	}

	if req.ActionType == 1 { // Favorite action
		if flag { // If favorite record already exists, do nothing
			return nil
		}

		// Create the record
		err := db.CreateFavoriteRecord(s.ctx, req.UserId, req.VideoId)
		if err != nil {
			return err
		}
	} else if req.ActionType == 2 { // Cancel favorite action
		if !flag { // If favorite record already not exists, do nothing
			return nil
		}

		// Delete the record
		err := db.DeleteFavoriteRecord(s.ctx, req.UserId, req.VideoId)
		if err != nil {
			return err
		}
	}

	// Do rpc call to add or subtract favorite count for the video
	resp, err := rpc.UpdateFavoriteCount(s.ctx, &video.UpdateFavoriteCountRequest{
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
