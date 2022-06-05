package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/CharmingCharm/DouSheng/internal/api/rpc"

	"github.com/CharmingCharm/DouSheng/kitex_gen/user"
	"github.com/CharmingCharm/DouSheng/pkg/constants"
	"github.com/CharmingCharm/DouSheng/pkg/middleware"
	"github.com/CharmingCharm/DouSheng/pkg/send"
	"github.com/CharmingCharm/DouSheng/pkg/status"
	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	constants.Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUIdString := c.Query("to_user_id")
	actionTypeString := c.Query("action_type")

	res := constants.Response{}

	if token == "" || toUIdString == "" || actionTypeString == "" {
		send.SendStatus(c, status.ParamErr, &res)
		return
	}

	toUId, err := strconv.ParseInt(toUIdString, 10, 64)
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}

	actionType64, err := strconv.ParseInt(actionTypeString, 10, 32)
	fmt.Println(actionType64)
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

	// rpc.action.UpdateRelationship()

	resp, err := rpc.UpdateUserFollow(context.Background(), &user.UpdateUserFollowRequest{
		UserId:     myId,
		ToUserId:   toUId,
		ActionType: actionType,
	})
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}

	send.SendResp(c, *resp.BaseResp, &res)
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: constants.Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: constants.Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}
