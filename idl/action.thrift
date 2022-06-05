namespace go action

include 'base.thrift'

struct UpdateFavoriteRequest {
    1:required i64  user_id
    2:required i64  video_id
    3:required i32  action_type
}

struct UpdateFavoriteResponse {
    1:required base.BaseResp     base_resp
}

struct GetFavoriteVideosRequest {
    1:required i64   user_id
    2:required i64   my_id
}

struct GetFavoriteVideosResponse {
    1:required base.BaseResp     base_resp
    2:required list<base.Video>  video_list
}

struct UpdateCommentRequest {
    1:required i64      user_id
    2:required i64      video_id
    3:required i32      action_type
    4:optional string   comment_text
    5:optional i64      comment_id
}

struct UpdateCommentResponse {
    1:required base.BaseResp base_resp
}

struct GetCommentListsRequest {
    1:required i64  user_id
    2:required i64  my_id
    3:required i64  video_id
}

struct GetCommentListsResponse {
    1:required base.BaseResp     base_resp
    2:required list<base.Comment>  comment_list
}

struct UpdateRelationshipRequest {
    1:required i64  user_id
    2:required i64  to_user_id
    3:required i32  action_type
}

struct UpdateRelationshipResponse {
    1:required base.BaseResp base_resp
}

struct GetUserFollowListRequest {
    1:required i64  user_id
    2:required i64  my_id
}

struct GetUserFollowListResponse {
    1:required base.BaseResp     base_resp
    2:required list<base.User>   user_list
}

struct GetUserFollowerListRequest {
    1:required i64  user_id
    2:required i64  my_id
}

struct GetUserFollowerListResponse {
    1:required base.BaseResp     base_resp
    2:required list<base.User>   user_list
}

struct CheckRelationRequest {
    1:required i64 my_id
    2:required i64 u_id
}

struct CheckRelationResponse {
    1:required base.BaseResp base_resp
    2:required bool     is_follow
}

service ActionService {
    UpdateFavoriteResponse UpdateFavorite(UpdateFavoriteRequest req)
    GetFavoriteVideosResponse GetFavoriteVideos(GetFavoriteVideosRequest req)
    UpdateCommentResponse UpdateComment(UpdateCommentRequest req)
    GetCommentListsResponse GetCommentLists(GetCommentListsRequest req)
    UpdateRelationshipResponse UpdateRelationship(UpdateRelationshipRequest req)
    GetUserFollowListResponse GetUserFollowList(GetUserFollowListRequest req)
    GetUserFollowerListResponse GetUserFollowerList(GetUserFollowerListRequest req)
    CheckRelationResponse CheckRelation(CheckRelationRequest req)
}