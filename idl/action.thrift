namespace go action

struct BaseResp {
    1:required i64       status_code
    2:optional string    status_message
    3:required i64       service_time
}

struct User {
    1:required i64       id
    2:required string    name
    3:optional i64       follow_count
    4:optional i64       follower_count
    5:required bool      is_follow
}

struct Video {
    1:required i64       id
    2:required User      author
    3:required string    play_url
    4:required string    cover_url
    5:required i64       favorite_count
    6:required i64       comment_count
    7:required bool      is_favorite
    8:required string    title
}

struct Comment {
    1:required i64      id
    2:required User     user
    3:required string   content
    4:required string   create_date
}

struct UpdateFavoriteRequest {
    1:required i64  user_id
    2:required i64  video_id
    3:required i32  action_type
}

struct UpdateFavoriteResponse {
    1:required BaseResp     base_resp
}

struct GetFavoriteVideosRequest {
    1:required i64   user_id
    2:optional i64   my_id
}

struct GetFavoriteVideosResponse {
    1:required BaseResp     base_resp
    2:required list<Video>  video_list
}

struct UpdateCommentRequest {
    1:required i64      user_id
    2:required i64      video_id
    3:required i32      action_type
    4:optional string   comment_text
    5:optional i64      comment_id
}

struct UpdateCommentResponse {
    1:required BaseResp base_resp
}

struct GetCommentListsRequest {
    1:required i64  user_id
    2:optional i64  my_id
    3:required i64  video_id
}

struct GetCommentListsResponse {
    1:required BaseResp     base_resp
    2:required list<Comment>  comment_list
}

struct UpdateRelationshipRequest {
    1:required i64  user_id
    2:required i64  to_user_id
    3:required i32  action_type
}

struct UpdateRelationshipResponse {
    1:required BaseResp base_resp
}

struct GetUserFollowListRequest {
    1:required i64  user_id
    2:optional i64  my_id
}

struct GetUserFollowListResponse {
    1:required BaseResp     base_resp
    2:required list<User>   user_list
}

struct GetUserFollowerListRequest {
    1:required i64  user_id
    2:optional i64  my_id
}

struct GetUserFollowerListResponse {
    1:required BaseResp     base_resp
    2:required list<User>   user_list
}

struct CheckRelationRequest {
    1:required i64 my_id
    2:required i64 u_id
}

struct CheckRelationResponse {
    1:required BaseResp base_resp
    2:required bool     is_follow
}

service ActionService {
    UpdateFavoriteResponse updateFavorite(UpdateFavoriteRequest req)
    GetFavoriteVideosResponse getFavoriteVideos(GetFavoriteVideosRequest req)
    UpdateCommentResponse updateComment(UpdateCommentRequest req)
    GetCommentListsResponse getCommentLists(GetCommentListsRequest req)
    UpdateRelationshipResponse updateRelationship(UpdateRelationshipRequest req)
    GetUserFollowListResponse getUserFollowList(GetUserFollowListRequest req)
    GetUserFollowerListResponse getUserFollowerList(GetUserFollowerListRequest req)
    CheckRelationResponse CheckRelation(CheckRelationRequest req)
}