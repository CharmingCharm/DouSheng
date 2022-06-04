package controller

import (
	"context"
	"net/http"

	"github.com/CharmingCharm/DouSheng/internal/api/rpc"
	"github.com/CharmingCharm/DouSheng/kitex_gen/user"
	"github.com/CharmingCharm/DouSheng/pkg/constants"
	"github.com/CharmingCharm/DouSheng/pkg/middleware"
	"github.com/CharmingCharm/DouSheng/pkg/status"

	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin

type UserRegisterResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	Token      string `json:"token"`
	UserID     int64  `json:"user_id"`
}

type UserLoginResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	Token      string `json:"token"`
	UserID     int64  `json:"user_id"`
}

// type UserInfoResponse struct {
// 	StatusCode int64     `json:"status_code"`
// 	StatusMsg  string    `json:"status_msg"`
// 	User       user.User `json:"user"`
// }

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	res := UserRegisterResponse{
		StatusCode: constants.DefaultStatusCode,
		StatusMsg:  constants.DefaultStatusMsg,
		UserID:     constants.DefaultErrPosInt64,
		Token:      constants.DefaultErrString,
	}

	if len(username) == 0 || len(password) == 0 {
		res.StatusCode = status.ParamErr.StatusCode
		res.StatusMsg = status.ParamErr.StatusMsg
		c.JSON(http.StatusOK, res)
	}

	resp, err := rpc.CreateUser(context.Background(), &user.CreateUserRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		st := status.ConvertErrorToStatus(err)
		res.StatusCode = st.StatusCode
		res.StatusMsg = st.StatusMsg
		c.JSON(http.StatusOK, res)
	}

	token, err := middleware.GenToken(username, resp.UserId)

	if err != nil {
		st := status.ConvertErrorToStatus(err)
		res.StatusCode = st.StatusCode
		res.StatusMsg = st.StatusMsg
		c.JSON(http.StatusOK, res)
		return
	}

	res.StatusCode = resp.BaseResp.StatusCode
	res.StatusMsg = resp.BaseResp.StatusMessage
	if res.StatusCode != status.SuccessCode {
		c.JSON(http.StatusOK, res)
		return
	}
	res.UserID = resp.UserId
	res.Token = token
	c.JSON(http.StatusOK, res)

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
		StatusCode: constants.DefaultStatusCode,
		StatusMsg:  constants.DefaultStatusMsg,
		UserID:     constants.DefaultErrPosInt64,
		Token:      constants.DefaultErrString,
	}

	if len(username) == 0 || len(password) == 0 {
		res.StatusCode = status.ParamErr.StatusCode
		res.StatusMsg = status.ParamErr.StatusMsg
		c.JSON(http.StatusOK, res)
		return
	}

	resp, err := rpc.CheckUser(context.Background(), &user.CheckUserRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		st := status.ConvertErrorToStatus(err)
		res.StatusCode = st.StatusCode
		res.StatusMsg = st.StatusMsg
		c.JSON(http.StatusOK, res)
		return
	}

	token, err := middleware.GenToken(username, resp.UserId)
	if err != nil {
		st := status.ConvertErrorToStatus(err)
		res.StatusCode = st.StatusCode
		res.StatusMsg = st.StatusMsg
		c.JSON(http.StatusOK, res)
		return
	}
	res.StatusCode = resp.BaseResp.StatusCode
	res.StatusMsg = resp.BaseResp.StatusMessage
	if res.StatusCode != status.SuccessCode {
		c.JSON(http.StatusOK, res)
		return
	}
	res.UserID = resp.UserId
	res.Token = token
	c.JSON(http.StatusOK, res)

}

func UserInfo(c *gin.Context) {
	// token := c.Query("token")

	// res := UserInfoResponse{
	// 	StatusCode: constants.DefaultStatusCode,
	// 	StatusMsg: constants.DefaultStatusMsg,
	// 	// User:	constants.D
	// }

	// claims, err := middleware.ParseToken(token)
	// if err != nil {

	// }
	// myId = claims.Id
	// fmt.Println(myId)

	// if user, exist := usersLoginInfo[token]; exist {
	// 	c.JSON(http.StatusOK, UserResponse{
	// 		Response: Response{StatusCode: 0},
	// 		User:     user,
	// 	})
	// } else {
	c.JSON(http.StatusOK, nil)
	// }
}
