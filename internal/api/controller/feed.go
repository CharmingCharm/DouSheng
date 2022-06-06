package controller

import (
	"context"
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
	// my_id := checkToken(token)
	// var Vs []Video
	// Vs, lastTime, err = rpc.LoadVideos(context.Background(), &video.LoadVideosRequest {
	// 	LastTime: lastTime,
	// 	MyId: my_id
	// })
	res := FeedResponse{}

	token := c.Query("token")
	lastTimeString := c.Query("last_time")

	req := video.LoadVideosRequest{
		MyId: -1,
	}

	if len(lastTimeString) != 0 {
		lastTime, err := strconv.ParseInt(lastTimeString, 10, 64)
		if err != nil {
			send.SendStatus(c, err, &res)
			return
		}
		req.LastTime = &lastTime
	}

	if len(token) != 0 {
		claims, err := middleware.ParseToken(token)
		if err != nil {
			send.SendStatus(c, err, &res)
			return
		}
		req.MyId = claims.Id
	}

	resp, err := rpc.LoadVideos(context.Background(), &req)

	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}

	res.VideoList = resp.VideoList
	if resp.NextTime != nil {
		res.NextTime = *resp.NextTime
	}
	send.SendResp(c, *resp.BaseResp, &res)
}
