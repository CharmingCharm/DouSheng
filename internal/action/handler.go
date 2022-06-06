package main

import (
	"context"

	action "github.com/CharmingCharm/DouSheng/kitex_gen/action"

	"github.com/CharmingCharm/DouSheng/internal/action/service"
	"github.com/CharmingCharm/DouSheng/pkg/response"
	"github.com/CharmingCharm/DouSheng/pkg/status"
)

// ActionServiceImpl implements the last service interface defined in the IDL.
type ActionServiceImpl struct{}

// UpdateFavorite implements the ActionServiceImpl interface.
func (s *ActionServiceImpl) UpdateFavorite(ctx context.Context, req *action.UpdateFavoriteRequest) (resp *action.UpdateFavoriteResponse, err error) {
	// TODO: Your code here...
	resp = new(action.UpdateFavoriteResponse)

	// type UpdateFavoriteRequest struct {
	// 	UserId     int64 `thrift:"user_id,1,required" json:"user_id"`
	// 	VideoId    int64 `thrift:"video_id,2,required" json:"video_id"`
	// 	ActionType int32 `thrift:"action_type,3,required" json:"action_type"`
	// }

	if req.VideoId <= 0 || req.UserId <= 0 || (req.ActionType != 1 && req.ActionType != 2) {
		resp.BaseResp = response.BuildBaseResp(status.ParamErr)
		return resp, nil
	}

	err = service.NewUpdateFavoriteService(ctx).UpdateFavorite(req)
	if err != nil {
		resp.BaseResp = response.BuildBaseResp(status.ConvertErrorToStatus(err))
		return resp, nil
	}
	resp.BaseResp = response.BuildBaseResp(status.Success)
	return resp, nil
}

// GetFavoriteVideos implements the ActionServiceImpl interface.
func (s *ActionServiceImpl) GetFavoriteVideos(ctx context.Context, req *action.GetFavoriteVideosRequest) (resp *action.GetFavoriteVideosResponse, err error) {
	// TODO: Your code here...
	resp = new(action.GetFavoriteVideosResponse)

	// type GetFavoriteVideosRequest struct {
	// 	UserId int64 `thrift:"user_id,1,required" json:"user_id"`
	// 	MyId   int64 `thrift:"my_id,2,required" json:"my_id"`
	// }

	if req.UserId <= 0 {
		resp.BaseResp = response.BuildBaseResp(status.ParamErr)
		return resp, nil
	}

	videoList, err := service.NewGetFavoriteVideosService(ctx).GetFavoriteVideos(req)
	if err != nil {
		resp.BaseResp = response.BuildBaseResp(status.ConvertErrorToStatus(err))
		return resp, nil
	}
	resp.BaseResp = response.BuildBaseResp(status.Success)
	resp.VideoList = videoList
	return resp, nil
}

// UpdateComment implements the ActionServiceImpl interface.
func (s *ActionServiceImpl) UpdateComment(ctx context.Context, req *action.UpdateCommentRequest) (resp *action.UpdateCommentResponse, err error) {
	// TODO: Your code here...
	resp = new(action.UpdateCommentResponse)

	// type UpdateCommentRequest struct {
	// 	UserId      int64   `thrift:"user_id,1,required" json:"user_id"`
	// 	VideoId     int64   `thrift:"video_id,2,required" json:"video_id"`
	// 	ActionType  int32   `thrift:"action_type,3,required" json:"action_type"`
	// 	CommentText *string `thrift:"comment_text,4" json:"comment_text,omitempty"`
	// 	CommentId   *int64  `thrift:"comment_id,5" json:"comment_id,omitempty"`
	// }

	if req.UserId <= 0 || req.VideoId <= 0 || (req.ActionType != 1 && req.ActionType != 2) {
		resp.BaseResp = response.BuildBaseResp(status.ParamErr)
		return resp, nil
	}

	comment, err := service.NewUpdateCommentService(ctx).UpdateComment(req)
	if err != nil {
		resp.BaseResp = response.BuildBaseResp(status.ConvertErrorToStatus(err))
		return resp, nil
	}
	resp.BaseResp = response.BuildBaseResp(status.Success)
	resp.Comment = comment
	return resp, nil
}

// GetCommentList implements the ActionServiceImpl interface.
func (s *ActionServiceImpl) GetCommentList(ctx context.Context, req *action.GetCommentListRequest) (resp *action.GetCommentListResponse, err error) {
	// TODO: Your code here...
	resp = new(action.GetCommentListResponse)

	// type GetCommentListRequest struct {
	// 	UserId  int64 `thrift:"user_id,1,required" json:"user_id"`
	// 	MyId    int64 `thrift:"my_id,2,required" json:"my_id"`
	// 	VideoId int64 `thrift:"video_id,3,required" json:"video_id"`
	// }

	if req.MyId <= 0 || req.VideoId <= 0 {
		resp.BaseResp = response.BuildBaseResp(status.ParamErr)
		return resp, nil
	}

	commentList, err := service.NewGetCommentListService(ctx).GetCommentList(req)
	if err != nil {
		resp.BaseResp = response.BuildBaseResp(status.ConvertErrorToStatus(err))
		return resp, nil
	}
	resp.BaseResp = response.BuildBaseResp(status.Success)
	resp.CommentList = commentList
	return resp, nil
}

// UpdateRelationship implements the ActionServiceImpl interface.
func (s *ActionServiceImpl) UpdateRelationship(ctx context.Context, req *action.UpdateRelationshipRequest) (resp *action.UpdateRelationshipResponse, err error) {
	// TODO: Your code here...
	resp = new(action.UpdateRelationshipResponse)

	// type UpdateRelationshipRequest struct {
	// 	UserId     int64 `thrift:"user_id,1,required" json:"user_id"`
	// 	ToUserId   int64 `thrift:"to_user_id,2,required" json:"to_user_id"`
	// 	ActionType int32 `thrift:"action_type,3,required" json:"action_type"`
	// }

	if req.UserId <= 0 || req.ToUserId <= 0 || (req.ActionType != 1 && req.ActionType != 2) {
		resp.BaseResp = response.BuildBaseResp(status.ParamErr)
		return resp, nil
	}

	err = service.NewUpdateRelationshipService(ctx).UpdateRelationship(req)
	if err != nil {
		resp.BaseResp = response.BuildBaseResp(status.ConvertErrorToStatus(err))
		return resp, nil
	}
	resp.BaseResp = response.BuildBaseResp(status.Success)
	return resp, nil
}

// GetUserFollowList implements the ActionServiceImpl interface.
func (s *ActionServiceImpl) GetUserFollowList(ctx context.Context, req *action.GetUserFollowListRequest) (resp *action.GetUserFollowListResponse, err error) {
	// TODO: Your code here...
	resp = new(action.GetUserFollowListResponse)

	// type GetUserFollowListRequest struct {
	// 	UserId int64 `thrift:"user_id,1,required" json:"user_id"`
	// 	MyId   int64 `thrift:"my_id,2,required" json:"my_id"`
	// }

	if req.UserId <= 0 {
		resp.BaseResp = response.BuildBaseResp(status.ParamErr)
		return resp, nil
	}

	userList, err := service.NewGetUserFollowListService(ctx).GetUserFollowList(req)
	if err != nil {
		resp.BaseResp = response.BuildBaseResp(status.ConvertErrorToStatus(err))
		return resp, nil
	}
	resp.BaseResp = response.BuildBaseResp(status.Success)
	resp.UserList = userList
	return resp, nil
}

// GetUserFollowerList implements the ActionServiceImpl interface.
func (s *ActionServiceImpl) GetUserFollowerList(ctx context.Context, req *action.GetUserFollowerListRequest) (resp *action.GetUserFollowerListResponse, err error) {
	// TODO: Your code here...
	resp = new(action.GetUserFollowerListResponse)

	// type GetUserFollowerListRequest struct {
	// 	UserId int64 `thrift:"user_id,1,required" json:"user_id"`
	// 	MyId   int64 `thrift:"my_id,2,required" json:"my_id"`
	// }

	if req.UserId <= 0 {
		resp.BaseResp = response.BuildBaseResp(status.ParamErr)
		return resp, nil
	}

	userList, err := service.NewGetUserFollowerListService(ctx).GetUserFollowerList(req)
	if err != nil {
		resp.BaseResp = response.BuildBaseResp(status.ConvertErrorToStatus(err))
		return resp, nil
	}
	resp.BaseResp = response.BuildBaseResp(status.Success)
	resp.UserList = userList
	return resp, nil
}

// CheckRelation implements the ActionServiceImpl interface.
func (s *ActionServiceImpl) CheckRelation(ctx context.Context, req *action.CheckRelationRequest) (resp *action.CheckRelationResponse, err error) {
	// TODO: Your code here...
	resp = new(action.CheckRelationResponse)

	// type CheckRelationRequest struct {
	// 	MyId   int64 `thrift:"my_id,1,required" json:"my_id"`
	// 	UserId int64 `thrift:"user_id,2,required" json:"user_id"`
	// }

	if req.MyId <= 0 || req.UserId <= 0 {
		resp.BaseResp = response.BuildBaseResp(status.ParamErr)
		return resp, nil
	}

	flag, err := service.NewCheckRelationService(ctx).CheckRelation(req)
	if err != nil {
		resp.BaseResp = response.BuildBaseResp(status.ConvertErrorToStatus(err))
		return resp, nil
	}
	resp.BaseResp = response.BuildBaseResp(status.Success)
	resp.IsFollow = &flag
	return resp, nil
}

// CheckFavorite implements the ActionServiceImpl interface.
func (s *ActionServiceImpl) CheckFavorite(ctx context.Context, req *action.CheckFavoriteRequest) (resp *action.CheckFavoriteResponse, err error) {
	// TODO: Your code here...

	// type CheckFavoriteRequest struct {
	// 	MyId    int64 `thrift:"my_id,1,required" json:"my_id"`
	// 	VideoId int64 `thrift:"video_id,2,required" json:"video_id"`
	// }
	resp = new(action.CheckFavoriteResponse)

	if req.MyId <= 0 || req.VideoId <= 0 {
		resp.BaseResp = response.BuildBaseResp(status.ParamErr)
		return resp, nil
	}

	flag, err := service.NewCheckFavoriteService(ctx).CheckFavorite(req)
	if err != nil {
		resp.BaseResp = response.BuildBaseResp(status.ConvertErrorToStatus(err))
		return resp, nil
	}
	resp.BaseResp = response.BuildBaseResp(status.Success)
	resp.IsFavorite = &flag
	return resp, nil
}
