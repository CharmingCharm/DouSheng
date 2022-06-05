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
	Token  string `json:"token"`
	UserID int64  `json:"user_id"`
}

type UserLoginResponse struct {
	constants.Response
	Token  string `json:"token"`
	UserID int64  `json:"user_id"`
}

type UserInfoResponse struct {
	constants.Response
	User base.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	res := UserRegisterResponse{
		UserID: constants.DefaultErrPosInt64,
		Token:  constants.DefaultErrString,
	}

	if len(username) == 0 || len(password) == 0 {
		send.SendStatus(c, status.ParamErr, &res)
		return
	}

	resp, err := rpc.CreateUser(context.Background(), &user.CreateUserRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}

	token, err := middleware.GenToken(username, resp.UserId)

	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}

	if resp.BaseResp.StatusCode != status.SuccessCode {
		send.SendResp(c, *resp.BaseResp, &res)
		return
	}
	res.UserID = resp.UserId
	res.Token = token
	send.SendResp(c, *resp.BaseResp, &res)

	// if _, exist := usersLoginInfo[token]; exist {
	// 	c.JSON(http.StatusOK, UserLoginResponse{
	// 		Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
	// 	})
	// } else {
	// 	atomic.AddInt64(&userIdSequence, 1)
	// 	newUser := User{
	// 		Id:   userIdSequence,
	// 		Name: username,
	// 	}
	// 	usersLoginInfo[token] = newUser
	// 	c.JSON(http.StatusOK, UserLoginResponse{
	// 		Response: Response{StatusCode: 0},
	// 		UserId:   userIdSequence,
	// 		Token:    username + password,
	// 	})
	// }

	// var registerVar UserParam
	// if err := c.ShouldBind(&registerVar); err != nil {
	// 	SendResponse(c, errno.ConvertErr(err), nil)
	// 	return
	// }

	// if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
	// 	SendResponse(c, errno.ParamErr, nil)
	// 	return
	// }

	// err := rpc.CreateUser(context.Background(), &userdemo.CreateUserRequest{
	// 	UserName: registerVar.UserName,
	// 	Password: registerVar.PassWord,
	// })
	// if err != nil {
	// 	SendResponse(c, errno.ConvertErr(err), nil)
	// 	return
	// }
	// SendResponse(c, errno.Success, nil)
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	res := UserLoginResponse{
		UserID: constants.DefaultErrPosInt64,
		Token:  constants.DefaultErrString,
	}
	res.StatusCode = constants.DefaultStatusCode
	res.StatusMsg = constants.DefaultStatusMsg

	if len(username) == 0 || len(password) == 0 {
		send.SendStatus(c, status.ParamErr, &res)
		return
	}

	resp, err := rpc.CheckUser(context.Background(), &user.CheckUserRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}

	token, err := middleware.GenToken(username, resp.UserId)
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}

	if resp.BaseResp.StatusCode != status.SuccessCode {
		send.SendResp(c, *resp.BaseResp, &res)
		return
	}
	res.UserID = resp.UserId
	res.Token = token
	send.SendResp(c, *resp.BaseResp, &res)
}

func UserInfo(c *gin.Context) {

	res := UserInfoResponse{
		User: base.User{
			Id:            constants.DefaultErrPosInt64,
			Name:          constants.DefaultErrString,
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		},
	}
	res.StatusCode = constants.DefaultStatusCode
	res.StatusMsg = constants.DefaultStatusMsg

	uIdInString := c.Query("user_id")
	token := c.Query("token")

	if uIdInString == "" || token == "" {
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
	if myId == uId {
		myId = -1
	}

	resp, err := rpc.GetUserInfo(context.Background(), &user.GetUserInfoRequest{
		UserId: uId,
		MyId:   myId,
	})

	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}

	if resp.BaseResp.StatusCode != status.SuccessCode {
		send.SendResp(c, *resp.BaseResp, &res)
		return
	}
	res.User = *resp.User
	send.SendResp(c, *resp.BaseResp, &res)
}
