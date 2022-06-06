package controller

import (
	"context"
	"strconv"

	"github.com/CharmingCharm/DouSheng/kitex_gen/action"
	"github.com/CharmingCharm/DouSheng/kitex_gen/base"

	"github.com/CharmingCharm/DouSheng/internal/api/rpc"

	"github.com/CharmingCharm/DouSheng/pkg/constants"
	"github.com/CharmingCharm/DouSheng/pkg/middleware"
	"github.com/CharmingCharm/DouSheng/pkg/send"
	"github.com/CharmingCharm/DouSheng/pkg/status"
	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	res := constants.Response{}

	// token := c.Query("token")
	// uIdInString := c.Query("user_id")

	token := c.Query("token")
	videoIdString := c.Query("video_id")
	actionTypeString := c.Query("action_type")

	if len(token) == 0 || len(videoIdString) == 0 || len(actionTypeString) == 0 {
		send.SendStatus(c, status.ParamErr, &res)
		return
	}

	claims, err := middleware.ParseToken(token)
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}
	myId := claims.Id

	vId, err := strconv.ParseInt(videoIdString, 10, 64)
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}

	actionType64, err := strconv.ParseInt(actionTypeString, 10, 32)
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}
	actionType := int32(actionType64)
	if actionType != 1 && actionType != 2 {
		send.SendStatus(c, status.ParamErr, &res)
		return
	}

	resp, err := rpc.UpdateFavorite(context.Background(), &action.UpdateFavoriteRequest{
		UserId:     myId,
		VideoId:    vId,
		ActionType: actionType,
	})

	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}
	send.SendResp(c, *resp.BaseResp, &res)
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {

	res := VideoListResponse{
		VideoList: make([]*base.Video, 0),
	}

	token := c.Query("token")
	uIdInString := c.Query("user_id")

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
	if myId == uId {
		myId = -1
	}

	resp, err := rpc.GetFavoriteVideos(context.Background(), &action.GetFavoriteVideosRequest{
		UserId: uId,
		MyId:   myId,
	})
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}
	res.VideoList = resp.VideoList
	send.SendResp(c, *resp.BaseResp, &res)
	// c.JSON(http.StatusOK, VideoListResponse{
	// 	Response: Response{
	// 		StatusCode: 0,
	// 	},
	// 	VideoList: DemoVideos,
	// })
}
