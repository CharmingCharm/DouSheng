package controller

import (
	"context"
	"fmt"
	"strconv"

	"github.com/CharmingCharm/DouSheng/pkg/constants"
	"github.com/CharmingCharm/DouSheng/pkg/send"

	"github.com/CharmingCharm/DouSheng/internal/api/rpc"
	"github.com/CharmingCharm/DouSheng/kitex_gen/base"
	"github.com/CharmingCharm/DouSheng/kitex_gen/video"
	"github.com/CharmingCharm/DouSheng/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	constants.Response
	VideoList []*base.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	// Initialize the response object
	res := FeedResponse{}

	// Do the argument checking and generate rpc request object
	token := c.Query("token")
	lastTimeString := c.Query("last_time")

	req := video.LoadVideosRequest{}

	if len(lastTimeString) != 0 {
		lastTime, err := strconv.ParseInt(lastTimeString, 10, 64)
		if err != nil {
			send.SendStatus(c, err, &res)
			return
		}
		req.LastTime = &lastTime
	}

	req.MyId = -1
	if len(token) != 0 {
		claims, err := middleware.ParseToken(token)
		if err != nil {
			send.SendStatus(c, err, &res)
			return
		}
		req.MyId = claims.Id
	}

	// Fetch videos and handle rpc error
	resp, err := rpc.LoadVideos(context.Background(), &req)
	if err != nil {
		fmt.Println()
		send.SendStatus(c, err, &res)
		return
	}

	// Respoonse
	res.VideoList = resp.VideoList
	if resp.NextTime != nil {
		res.NextTime = *resp.NextTime
	}
	send.SendResp(c, *resp.BaseResp, &res)
}
