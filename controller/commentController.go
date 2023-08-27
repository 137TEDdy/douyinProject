/*
   @Author Ted
   @Since 2023/8/2 9:07
*/

package controller

import (
	. "douyinProject/common"
	"douyinProject/log"
	"douyinProject/model"
	"douyinProject/service"
	"douyinProject/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 登录用户对视频进行评论
func CommentAction(c *gin.Context) {
	var comment_text, comment_id string
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")
	userTmp, _ := c.Get("user") //中间件处理后获取user
	user := userTmp.(model.User)
	actiontype, _ := strconv.ParseInt(action_type, 10, 64) //转成int64，  参数含义：十进制的64位
	videoid, _ := strconv.ParseInt(video_id, 10, 64)

	if actiontype == 1 {
		//发表评论：视频id，用户id，评论内容；插入comment表里； 并返回评论内容
		comment_text = c.Query("comment_text")
		comment, err := service.PublishComment(user.Id, videoid, comment_text, utils.GetCurrentTimeMMDD())
		if err != nil {
			log.Error(err.Error())
			Resp(c, CodeCommentError, Response{-1, Msg(CodeCommentError)})
			return
		}
		Resp(c, 200, CommentResponse{
			Response{0, Msg(CodeSuccess)},
			comment,
		})
		return

	} else if actiontype == 2 {
		//删除评论，根据评论id删除
		comment_id = c.Query("comment_id")
		commentId, err := strconv.ParseInt(comment_id, 10, 64)
		//调用service，删除评论
		err = service.DeleteComment(commentId)
		if err != nil {
			log.Error(err.Error())
			Resp(c, CodeCommentError, Response{-1, Msg(CodeCommentError)})
			return
		}
		Resp(c, CodeSuccess, Response{0, Msg(CodeSuccess)})
	}

}

// 获取评论列表,根据视频id查询所有评论
func GetCommentList(c *gin.Context) {
	video_id := c.Query("video_id")
	videoId, err := strconv.ParseInt(video_id, 10, 64)
	if err != nil {
		log.Error(err.Error())
		Resp(c, CodeInvalidParams, Response{-1, Msg(CodeInvalidParams)})
		return
	}

	commentList, err := service.GetCommentList(videoId)
	if err != nil {
		log.Error(err.Error())
		Resp(c, CodeCommentError, Response{-1, Msg(CodeCommentError)})
		return
	}
	Resp(c, CodeSuccess, CommentListResponse{
		Response{0, Msg(CodeSuccess)},
		commentList,
	})
}
