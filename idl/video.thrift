namespace go video

include 'base.thrift'

struct GetVideoListRequest {
    1:required list<i64>    video_ids
    2:required i64          user_id
}

struct GetVideoListResponse {
    1:required base.BaseResp     base_resp
    2:required list<base.Video>  videos
}

struct LoadVideosRequest {
    1:optional i64   last_time
    2:required i64   my_id
}

struct LoadVideosResponse {
    1:required base.BaseResp     base_resp
    2:required list<base.Video>  video_list
    3:optional i64          next_time
}

struct PublishVideoRequest {
    1:required i64      my_id
    2:required string   data_url
    3:required string   title
}

struct PublishVideoResponse {
    1:required base.BaseResp base_resp
}

struct GetPublishedVideosRequest {
    1:required i64  user_id
    2:required i64  my_id
}

struct GetPublishedVideosResponse {
    1:required base.BaseResp     base_resp
    2:required list<base.Video>  video_list
}

struct UpdateFavoriteCountRequest {
    1:required i64  video_id
    2:required i32  action_type
}

struct UpdateFavoriteCountResponse {
    1:required base.BaseResp base_resp
}

struct UpdateCommentCountRequest {
    1:required i64  video_id
    2:required i32  action_type
}

struct UpdateCommentCountResponse {
    1:required base.BaseResp base_resp
}

service VideoService {
    GetVideoListResponse GetVideoList(GetVideoListRequest req)
    LoadVideosResponse LoadVideos(LoadVideosRequest req)
    PublishVideoResponse PublishVideo(PublishVideoRequest req)
    GetPublishedVideosResponse GetPublishedVideos(GetPublishedVideosRequest req)
    UpdateFavoriteCountResponse UpdateFavoriteCount(UpdateFavoriteCountRequest req)
    UpdateCommentCountResponse UpdateCommentCount(UpdateCommentCountRequest req)
}