package main

import (
	"context"
	"github.com/CharmingCharm/DouSheng/idl/kitex_gen/action"
)

// ActionServiceImpl implements the last service interface defined in the IDL.
type ActionServiceImpl struct{}

// UpdateFavorite implements the ActionServiceImpl interface.
func (s *ActionServiceImpl) UpdateFavorite(ctx context.Context, req *action.UpdateFavoriteRequest) (resp *action.UpdateFavoriteResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFavoriteVideos implements the ActionServiceImpl interface.
func (s *ActionServiceImpl) GetFavoriteVideos(ctx context.Context, req *action.GetFavoriteVideosRequest) (resp *action.GetFavoriteVideosResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateComment implements the ActionServiceImpl interface.
func (s *ActionServiceImpl) UpdateComment(ctx context.Context, req *action.UpdateCommentRequest) (resp *action.UpdateCommentResponse, err error) {
	// TODO: Your code here...
	return
}

// GetCommentLists implements the ActionServiceImpl interface.
func (s *ActionServiceImpl) GetCommentLists(ctx context.Context, req *action.GetCommentListsRequest) (resp *action.GetCommentListsResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateRelationship implements the ActionServiceImpl interface.
func (s *ActionServiceImpl) UpdateRelationship(ctx context.Context, req *action.UpdateRelationshipRequest) (resp *action.UpdateRelationshipResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserFollowList implements the ActionServiceImpl interface.
func (s *ActionServiceImpl) GetUserFollowList(ctx context.Context, req *action.GetUserFollowListRequest) (resp *action.GetUserFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserFollowerList implements the ActionServiceImpl interface.
func (s *ActionServiceImpl) GetUserFollowerList(ctx context.Context, req *action.GetUserFollowerListRequest) (resp *action.GetUserFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}
