package controller

import (
	"context"
	"strconv"

	"github.com/CharmingCharm/DouSheng/internal/api/rpc"
	"github.com/CharmingCharm/DouSheng/kitex_gen/action"
	"github.com/CharmingCharm/DouSheng/kitex_gen/base"
	"github.com/CharmingCharm/DouSheng/pkg/constants"
	"github.com/CharmingCharm/DouSheng/pkg/middleware"
	"github.com/CharmingCharm/DouSheng/pkg/send"
	"github.com/CharmingCharm/DouSheng/pkg/status"
	"github.com/gin-gonic/gin"
)

type GetCommentListResponse struct {
	constants.Response
	CommentList []*base.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	constants.Response
	Comment *base.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	// Initialize response object
	res := CommentActionResponse{}

	// Do the argument checking
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

	commentReq := action.UpdateCommentRequest{
		UserId:     myId,
		VideoId:    vId,
		ActionType: actionType,
	}

	// Based on different operation, the arguments are different
	if actionType == 1 { // Add comment: check the comment_text and add it as rpc argument
		if len(c.Query("comment_text")) == 0 {
			send.SendStatus(c, status.ParamErr, &res)
			return
		}
		commentText := c.Query("comment_text")
		commentReq.CommentText = &commentText
	} else if actionType == 2 { // Delete comment: check comment_id and add it as rpc argument
		if len(c.Query("comment_id")) == 0 {
			send.SendStatus(c, status.ParamErr, &res)
			return
		}
		cId, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		if err != nil {
			send.SendStatus(c, err, &res)
			return
		}
		commentReq.CommentId = &cId
	}

	// Rpc call to update comment info
	resp, err := rpc.UpdateComment(context.Background(), &commentReq)
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}

	// Response
	if resp.Comment != nil {
		res.Comment = resp.Comment
	}
	send.SendResp(c, *resp.BaseResp, &res)
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	// Initialize response object
	res := GetCommentListResponse{}

	// Do argument checking
	token := c.Query("token")
	videoIdString := c.Query("video_id")

	if len(token) == 0 || len(videoIdString) == 0 {
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

	// Rpc call to get the comment list
	resp, err := rpc.GetCommentList(context.Background(), &action.GetCommentListRequest{
		MyId:    myId,
		VideoId: vId,
	})
	if err != nil {
		send.SendStatus(c, err, &res)
		return
	}

	// Response
	res.CommentList = resp.CommentList
	send.SendResp(c, *resp.BaseResp, &res)
}
