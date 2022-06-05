namespace go user

include 'base.thrift'

struct CreateUserRequest {
    1:required string    username
    2:required string    password
}

struct CreateUserResponse {
    1:required base.BaseResp  base_resp
    2:required i64       user_id
}

struct CheckUserRequest {
    1:required string    username
    2:required string    password
}

struct CheckUserResponse {
    1:required base.BaseResp  base_resp
    2:required i64       user_id
}

struct GetUserInfoRequest {
    1:required i64  user_id
    2:required i64  my_id
}

struct GetUserInfoResponse {
    1:required base.BaseResp  base_resp
    2:required base.User      user
}

struct UpdateUserFollowRequest {
    1:required i64       user_id
    2:required i64       to_user_id
    3:required i32       action_type
}

struct UpdateUserFollowResponse {
    1:required base.BaseResp  base_resp
}

service UserService {
    CreateUserResponse CreateUser(CreateUserRequest req)
    CheckUserResponse CheckUser(CheckUserRequest req)
    GetUserInfoResponse GetUserInfo(GetUserInfoRequest req)
    UpdateUserFollowResponse UpdateUserFollow(UpdateUserFollowRequest req)
}
