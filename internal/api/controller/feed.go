package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// var lastTime int64
// lastTime = -1

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	// my_id := checkToken(token)
	// var Vs []Video
	// Vs, lastTime, err = rpc.LoadVideos(context.Background(), &video.LoadVideosRequest {
	// 	LastTime: lastTime,
	// 	MyId: my_id
	// })
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
