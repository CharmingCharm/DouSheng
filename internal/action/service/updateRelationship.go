package service

import (
	"context"

	"github.com/CharmingCharm/DouSheng/internal/action/db"
	"github.com/CharmingCharm/DouSheng/internal/action/rpc"
	"github.com/CharmingCharm/DouSheng/pkg/status"

	"github.com/CharmingCharm/DouSheng/kitex_gen/action"
	"github.com/CharmingCharm/DouSheng/kitex_gen/user"
)

type UpdateRelationshipService struct {
	ctx context.Context
}

// NewUpdateRelationshipService new UpdateRelationshipService
func NewUpdateRelationshipService(ctx context.Context) *UpdateRelationshipService {
	return &UpdateRelationshipService{ctx: ctx}
}

// Update "Follow" related data
func (s *UpdateRelationshipService) UpdateRelationship(req *action.UpdateRelationshipRequest) error {
	// Check if the relation of user and to_user already exists.
	flag, err := db.FindRelationRecord(s.ctx, req.UserId, req.ToUserId)
	if err != nil {
		return err
	}

	if req.ActionType == 1 { // Add the relation ship
		if flag { // If relationship already exists, return nothing
			return nil
		}
		// Create the relationship record
		err := db.CreateRelationshipRecord(s.ctx, req.UserId, req.ToUserId)
		if err != nil {
			return err
		}
	} else if req.ActionType == 2 { // Cancel the relation ship
		if !flag { // If relationship not exists, return nothing
			return nil
		}
		// Delete the relationship record
		err := db.DeleteRelationshipRecord(s.ctx, req.UserId, req.ToUserId)
		if err != nil {
			return err
		}
	}

	// Do rpc call to udpate the user's follow count and follower count on user server
	resp, err := rpc.UpdateUserFollow(s.ctx, &user.UpdateUserFollowRequest{
		UserId:     req.UserId,
		ToUserId:   req.ToUserId,
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
