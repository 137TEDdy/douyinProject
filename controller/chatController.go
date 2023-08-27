package controller

import (
	. "douyinProject/common"
	"douyinProject/log"
	"douyinProject/model"
	"douyinProject/service"
	"douyinProject/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func ChatAction(c *gin.Context) {

	to_user_id := c.Query("to_user_id")
	action := c.Query("action_type")
	userTmp, _ := c.Get("user")
	content := c.Query("content")
	user := userTmp.(model.User)

	touserid, _ := strconv.ParseInt(to_user_id, 10, 64)
	message := model.Message{
		UserId:     user.Id,
		ToUserId:   touserid,
		ActionType: 0,
		Content:    content,
		CreateTime: utils.GetCurrentTimeForString(),
	}
	if action != "" {
		_, err := service.PublishMessage(message)
		if err != nil {
			log.Error(err.Error())
			Resp(c, CodeMessageError, Response{-1, Msg(CodeMessageError)})
			return
		}
		Resp(c, CodeSuccess, Response{0, Msg(CodeSuccess)})
		return
	}
}
func GetChatList(c *gin.Context) {
	to_user_id := c.Query("to_user_id")
	userTmp, _ := c.Get("user")
	user := userTmp.(model.User)
	user_id := user.Id

	touserid, _ := strconv.ParseInt(to_user_id, 10, 64)
	//获取当前时间，并格式化
	timeTmp := c.Query("pre_msg_time")
	timestamp, _ := strconv.ParseInt(timeTmp, 10, 64)
	// 根据秒级时间戳创建time.Time对象
	t := time.Unix(timestamp, 0)
	// 将time.Time对象格式化为年月日时分的字符串
	timeStr := t.Format("200601021504")
	log.Info("参数里的时间为:", timeStr)

	message, err := service.GetChatList(user_id, touserid, timeStr)
	Resp(c, CodeSuccess, MessageListResponse{
		Response: Response{0, Msg(CodeSuccess)},
		Messages: message,
	})
	if err != nil {
		log.Error("获取消息列表错误")
	}
	return
	//
	//log.Info(message)
	//ticker := time.NewTicker(2 * time.Second)
	//for range ticker.C {
	//	if err != nil {
	//		log.Error(err.Error())
	//		c.JSON(500, common.MessageListResponse{
	//			Response: common.Response{-1, "获取聊天记录失败"},
	//			Messages: message,
	//		})
	//		continue
	//	}
	//	c.JSON(200, common.MessageListResponse{
	//		Response: common.Response{0, "获取聊天记录"},
	//		Messages: message,
	//	})
	//	log.Info(message)
	//	continue
	//}

}
