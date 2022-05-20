namespace go user

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

struct CreateUserRequest {
    1:required string    username
    2:required string    password
}

struct CreateUserResponse {
    1:required BaseResp  base_resp
}

struct CheckUserRequest {
    1:required string    username
    2:required string    password
}

struct CheckUserResponse {
    1:required BaseResp  base_resp
}

struct GetUserInfoRequest {
    1:required i64  user_id
    2:optional i64  my_id
}

struct GetUserInfoResponse {
    1:required BaseResp  base_resp
    2:required User      user
}

struct UpdateUserFollowRequest {
    1:required i64       user_id
    2:required i64       to_user_id
    3:required i32       action_type
}

struct UpdateUserFollowResponse {
    1:required BaseResp  base_resp
}

service UserService {
    CreateUserResponse createUser(CreateUserRequest req)
    CheckUserResponse checkUser(CheckUserRequest req)
    GetUserInfoResponse getUserInfo(GetUserInfoRequest req)
    UpdateUserFollowResponse updateUserFollow(UpdateUserFollowRequest req)
}
