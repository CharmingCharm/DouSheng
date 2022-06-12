package controller

import (
	"context"
	"strconv"

	"github.com/CharmingCharm/DouSheng/internal/api/rpc"

	"github.com/CharmingCharm/DouSheng/kitex_gen/action"
	"github.com/CharmingCharm/DouSheng/kitex_gen/base"
	"github.com/CharmingCharm/DouSheng/pkg/constants"
	"github.com/CharmingCharm/DouSheng/pkg/middleware"
	"github.com/CharmingCharm/DouSheng/pkg/send"
	"github.com/CharmingCharm/DouSheng/pkg/status"
	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	constants.Response
	UserList []*base.User `json:"user_list,omitempty"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	// Initializing response object
	res := constants.Response{}

	// Do the argument checking and rpc arguments extracting
	token := c.Query("token")
	toUIdString := c.Query("to_user_id")
	actionTypeString := c.Query("action_type")

	if len(token) == 0 || len(toUIdString) == 0 || len(actionTypeString) == 0 {
		send.SendStatus(c, status.ParamErr, &res)
		return
	}

	toUId, err := strconv.ParseInt(toUIdString, 10, 64)
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}

	actionType64, err := strconv.ParseInt(actionTypeString, 10, 32)
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}
	if actionType64 != 1 && actionType64 != 2 {
		send.SendStatus(c, status.ParamErr, &res)
		return
	}
	actionType := int32(actionType64)

	claims, err := middleware.ParseToken(token)
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}
	myId := claims.Id

	// Call rpc function to do the relationship operation logic
	resp, err := rpc.UpdateRelationship(context.Background(), &action.UpdateRelationshipRequest{
		UserId:     myId,
		ToUserId:   toUId,
		ActionType: actionType,
	})
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}

	// Send back response
	send.SendResp(c, *resp.BaseResp, &res)
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	// Initialize response object
	res := UserListResponse{
		UserList: make([]*base.User, 0),
	}

	// Do the argument checking
	token := c.Query("token")
	uIdString := c.Query("user_id")

	if len(token) == 0 || len(uIdString) == 0 {
		send.SendStatus(c, status.ParamErr, &res)
		return
	}

	uId, err := strconv.ParseInt(uIdString, 10, 64)
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}

	claims, err := middleware.ParseToken(token)
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}
	myId := claims.Id

	// Call rpc function to fetch the user follow list
	resp, err := rpc.GetUserFollowList(context.Background(), &action.GetUserFollowListRequest{
		UserId: uId,
		MyId:   myId,
	})
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}

	// Fillin the response object and send back to client
	res.UserList = resp.UserList
	send.SendResp(c, *resp.BaseResp, &res)
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	// Initializing response object
	res := UserListResponse{
		UserList: make([]*base.User, 0),
	}

	// Do the argument checking and rpc arguments extracting
	token := c.Query("token")
	uIdString := c.Query("user_id")

	if len(token) == 0 || len(uIdString) == 0 {
		send.SendStatus(c, status.ParamErr, &res)
		return
	}

	uId, err := strconv.ParseInt(uIdString, 10, 64)
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}

	claims, err := middleware.ParseToken(token)
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}
	myId := claims.Id

	// Call rpc function to fetch the user follower list
	resp, err := rpc.GetUserFollowerList(context.Background(), &action.GetUserFollowerListRequest{
		UserId: uId,
		MyId:   myId,
	})
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}

	// Fillin the response object and send back to client
	res.UserList = resp.UserList
	send.SendResp(c, *resp.BaseResp, &res)
}
