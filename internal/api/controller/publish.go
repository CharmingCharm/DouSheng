package controller

import (
	"github.com/CharmingCharm/DouSheng/internal/api/ostg"
	"github.com/CharmingCharm/DouSheng/kitex_gen/video"
	"github.com/CharmingCharm/DouSheng/pkg/constants"
	"github.com/CharmingCharm/DouSheng/pkg/middleware"
	"github.com/CharmingCharm/DouSheng/pkg/send"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	res := video.PublishVideoResponse{}
	token := c.PostForm("token")
	claims, err := middleware.ParseToken(token)
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}

	myId := claims.Id

	// if _, exist := usersLoginInfo[token]; !exist {
	file, fileHeader, err := c.Request.FormFile("data")
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}
	// 上传文件到minio
	err = ostg.UploadVideo(fileHeader.Filename, fileHeader)

	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}
	playUrl := constants.MinIOEndpoint + "/video/" + fileHeader.Filename
	// TODO:调用RPC服务往video表里插入记录

	return
	// }

	// data, err := c.FormFile("data")
	// if err != nil {
	// 	c.JSON(http.StatusOK, Response{
	// 		StatusCode: 1,
	// 		StatusMsg:  err.Error(),
	// 	})
	// 	return
	// }

	// filename := filepath.Base(data.Filename)
	// user := usersLoginInfo[token]
	// finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	// saveFile := filepath.Join("./public/", finalName)
	// if err := c.SaveUploadedFile(data, saveFile); err != nil {
	// 	c.JSON(http.StatusOK, Response{
	// 		StatusCode: 1,
	// 		StatusMsg:  err.Error(),
	// 	})
	// 	return
	// }

	// c.JSON(http.StatusOK, Response{
	// 	StatusCode: 0,
	// 	StatusMsg:  finalName + " uploaded successfully",
	// })
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
