package controller

import (
	"context"
	"strconv"

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
	VideoList []*base.Video `json:"video_list"`
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

	/************** Uncheck **************/
	// if _, exist := usersLoginInfo[token]; !exist {
	// file, fileHeader, err := c.Request.FormFile("data")
	_, _, err = c.Request.FormFile("data")
	// fmt.Println(file)
	// fmt.Println(fileHeader)
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}
	// // 上传文件到minio
	// err = ostg.UploadVideo(fileHeader.Filename, fileHeader)

	// if err != nil {
	// 	send.SendStatus(c, err, &res)
	// 	return
	// }
	// playUrl := constants.MinIOEndpoint + "/video/" + fileHeader.Filename
	// fmt.Println(playUrl)
	// TODO:调用RPC服务往video表里插入记录
	/************** Uncheck **************/

	resp, err := rpc.PublishVideo(context.Background(), &video.PublishVideoRequest{
		MyId:    myId,
		DataUrl: "https://www.w3schools.com/html/movie.mp4", // Uncheck
		Title:   title,
	})

	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}
	send.SendResp(c, *resp.BaseResp, &res)
	// return
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
	// 	StatusMsg:  fileHeader.Filename + " uploaded successfully",
	// })
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	res := VideoListResponse{
		VideoList: nil,
	}

	token := c.Query("token")
	uIdInString := c.Query("user_id")

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

	resp, err := rpc.GetPublishedVideos(context.Background(), &video.GetPublishedVideosRequest{
		UserId: uId,
		MyId:   &myId,
	})

	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}

	if resp.BaseResp.StatusCode != status.SuccessCode {
		send.SendResp(c, *resp.BaseResp, &res)
		return
	}
	res.VideoList = resp.VideoList
	send.SendResp(c, *resp.BaseResp, &res)
}
