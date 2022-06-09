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

// CreateUser create user info.
func (s *UpdateRelationshipService) UpdateRelationship(req *action.UpdateRelationshipRequest) error {
	flag, err := db.FindRelationRecord(s.ctx, req.UserId, req.ToUserId)
	if err != nil {
		return err
	}

	if req.ActionType == 1 {
		if flag {
			return nil
		}
		err := db.CreateRelationshipRecord(s.ctx, req.UserId, req.ToUserId)
		if err != nil {
			return err
		}
	} else if req.ActionType == 2 {
		if !flag {
			return nil
		}
		err := db.DeleteRelationshipRecord(s.ctx, req.UserId, req.ToUserId)
		if err != nil {
			return err
		}
	}

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
