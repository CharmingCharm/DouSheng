package controller

import (
	"context"
	"strconv"

	"github.com/CharmingCharm/DouSheng/internal/api/rpc"
	"github.com/CharmingCharm/DouSheng/kitex_gen/base"
	"github.com/CharmingCharm/DouSheng/kitex_gen/user"
	"github.com/CharmingCharm/DouSheng/pkg/constants"
	"github.com/CharmingCharm/DouSheng/pkg/middleware"
	"github.com/CharmingCharm/DouSheng/pkg/send"
	"github.com/CharmingCharm/DouSheng/pkg/status"

	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin

type UserRegisterResponse struct {
	constants.Response
	Token  *string `json:"token,omitempty"`
	UserID *int64  `json:"user_id,omitempty"`
}

type UserLoginResponse struct {
	constants.Response
	Token  *string `json:"token,omitempty"`
	UserID *int64  `json:"user_id,omitempty"`
}

type UserInfoResponse struct {
	constants.Response
	User *base.User `json:"user,omitempty"`
}

func Register(c *gin.Context) {
	// Initialize response object
	res := UserRegisterResponse{}

	// Do the argument checking
	username := c.Query("username")
	password := c.Query("password")
	if len(username) == 0 || len(password) == 0 {
		send.SendStatus(c, status.ParamErr, &res)
		return
	}

	// Call rpc function to do the actual register logic
	resp, err := rpc.CreateUser(context.Background(), &user.CreateUserRequest{
		Username: username,
		Password: password,
	})

	// Error checking on rpc call
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}
	if resp.BaseResp.StatusCode != status.SuccessCode {
		send.SendResp(c, *resp.BaseResp, &res)
		return
	}

	// Fill in response object and return
	token, err := middleware.GenToken(username, *resp.UserId)
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}
	res.UserID = resp.UserId
	res.Token = &token
	send.SendResp(c, *resp.BaseResp, &res)
}

func Login(c *gin.Context) {
	// Initialize response object
	res := UserLoginResponse{}

	// Do the argument checking
	username := c.Query("username")
	password := c.Query("password")
	if len(username) == 0 || len(password) == 0 {
		send.SendStatus(c, status.ParamErr, &res)
		return
	}

	// Call rpc function to do the actual register logic
	resp, err := rpc.CheckUser(context.Background(), &user.CheckUserRequest{
		Username: username,
		Password: password,
	})

	// Error checking on rpc call
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}
	if resp.BaseResp.StatusCode != status.SuccessCode {
		send.SendResp(c, *resp.BaseResp, &res)
		return
	}

	// Fill in response object and return
	token, err := middleware.GenToken(username, *resp.UserId)
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}
	res.UserID = resp.UserId
	res.Token = &token
	send.SendResp(c, *resp.BaseResp, &res)
}

func UserInfo(c *gin.Context) {
	// Initialize response object
	res := UserInfoResponse{}

	// Do the argument checking and extracting
	uIdInString := c.Query("user_id")
	token := c.Query("token")
	if len(uIdInString) == 0 || len(token) == 0 {
		send.SendStatus(c, status.ParamErr, &res)
		return
	}

	uId, err := strconv.ParseInt(uIdInString, 10, 64)
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

	// Do rpc call to fetch the user information
	resp, err := rpc.GetUserInfo(context.Background(), &user.GetUserInfoRequest{
		UserId: uId,
		MyId:   myId,
	})

	// Error checking on rpc call
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}
	if resp.BaseResp.StatusCode != status.SuccessCode {
		send.SendResp(c, *resp.BaseResp, &res)
		return
	}

	// Fill in the response object for return
	res.User = resp.User
	send.SendResp(c, *resp.BaseResp, &res)
}
