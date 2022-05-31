package controller

import (
	"context"
	"net/http"

	"github.com/CharmingCharm/DouSheng/internal/api/rpc"
	"github.com/CharmingCharm/DouSheng/kitex_gen/user"
	"github.com/CharmingCharm/DouSheng/pkg/middleware"
	"github.com/CharmingCharm/DouSheng/pkg/status"
	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

type UserRegisterResponse struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	Token      string `json:"token"`       // 用户鉴权token
	UserID     int64  `json:"user_id"`     // 用户id
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	resp, err := rpc.CreateUser(context.Background(), &user.CreateUserRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		st := status.ConvertErrorToStatus(err)
		c.JSON(http.StatusOK, UserRegisterResponse{
			StatusCode: st.StatusCode,
			StatusMsg:  st.StatusMsg,
			Token:      "",
			UserID:     -1,
		})
	}

	token, err := middleware.GenToken(username, resp.UserId)

	if err != nil {
		st := status.ConvertErrorToStatus(err)
		c.JSON(http.StatusOK, UserRegisterResponse{
			StatusCode: st.StatusCode,
			StatusMsg:  st.StatusMsg,
			Token:      "",
			UserID:     -1,
		})
		return
	}

	if resp.BaseResp.StatusCode != status.SuccessCode {
		c.JSON(http.StatusOK, UserRegisterResponse{
			StatusCode: resp.BaseResp.StatusCode,
			StatusMsg:  resp.BaseResp.StatusMessage,
			Token:      "",
			UserID:     -1,
		})
		return
	}
	c.JSON(http.StatusOK, UserRegisterResponse{
		StatusCode: resp.BaseResp.StatusCode,
		StatusMsg:  resp.BaseResp.StatusMessage,
		Token:      token,
		UserID:     resp.UserId,
	})

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

	token := username + password

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
