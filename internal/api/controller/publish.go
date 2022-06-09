package controller

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/CharmingCharm/DouSheng/internal/api/ostg"

	"github.com/CharmingCharm/DouSheng/internal/api/rpc"

	"github.com/CharmingCharm/DouSheng/kitex_gen/base"
	"github.com/CharmingCharm/DouSheng/kitex_gen/video"
	"github.com/CharmingCharm/DouSheng/pkg/constants"
	"github.com/CharmingCharm/DouSheng/pkg/middleware"
	"github.com/CharmingCharm/DouSheng/pkg/send"
	"github.com/CharmingCharm/DouSheng/pkg/status"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	constants.Response
	VideoList []*base.Video `json:"video_list,omitempty"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	res := constants.Response{}

	token := c.PostForm("token")
	title := c.PostForm("title")

	claims, err := middleware.ParseToken(token)
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}

	if len(title) == 0 {
		title = constants.DefaultVideoTitle
	}

	myId := claims.Id
	username := claims.Username

	_, fileHeader, err := c.Request.FormFile("data")
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	split := strings.Split(fileHeader.Filename, ".")
	fileType := split[len(split)-1]
	err = ostg.UploadVideo(timestamp+"."+fileType, username, fileHeader)

	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}
	playUrl := constants.MinIOPos + "/video/" + timestamp + "." + fileType

	resp, err := rpc.PublishVideo(context.Background(), &video.PublishVideoRequest{
		MyId:    myId,
		DataUrl: playUrl, // Uncheck
		Title:   title,
	})
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}
	send.SendResp(c, *resp.BaseResp, &res)
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	res := VideoListResponse{}

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

	resp, err := rpc.GetPublishedVideos(context.Background(), &video.GetPublishedVideosRequest{
		UserId: uId,
		MyId:   myId,
	})

	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}
	res.VideoList = resp.VideoList
	send.SendResp(c, *resp.BaseResp, &res)
}
