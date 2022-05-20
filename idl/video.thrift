namespace go video

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

struct GetVideoListRequest {
    1:required list<i64>    video_ids
    2:optional i64          user_id
}

struct GetVideoListResponse {
    1:required BaseResp     base_resp
    2:required list<Video>  videos
}

struct LoadVideosRequest {
    1:optional i64   last_time
    2:optional i64   my_id
}

struct LoadVideosResponse {
    1:required BaseResp     base_resp
    2:required list<Video>  video_list
    3:optional i64          next_time
}

struct PublishVideoRequest {
    1:required i64      my_id
    2:required byte     data
    3:required string   title
}

struct PublishVideoResponse {
    1:required BaseResp base_resp
}

struct GetPublishedVideosRequest {
    1:required i64  user_id
    2:optional i64  my_id
}

struct GetPublishedVideosResponse {
    1:required BaseResp     base_resp
    2:required list<Video>  video_list
}

struct UpdateFavoriteCountRequest {
    1:required i64  video_id
    2:required i32  action_type
}

struct UpdateFavoriteCountResponse {
    1:required BaseResp base_resp
}

struct UpdateCommentCountRequest {
    1:required i64  video_id
    2:required i32  action_type
}

struct UpdateCommentCountResponse {
    1:required BaseResp base_resp
}

service VideoService {
    GetVideoListResponse getVideoList(GetVideoListRequest req)
    LoadVideosResponse loadVideos(LoadVideosRequest req)
    PublishVideoResponse publishVideo(PublishVideoRequest req)
    GetPublishedVideosResponse getPublishedVideos(GetPublishedVideosRequest req)
    UpdateFavoriteCountResponse updateFavoriteCount(UpdateFavoriteCountRequest req)
    UpdateCommentCountResponse updateCommentCount(UpdateCommentCountRequest req)
}