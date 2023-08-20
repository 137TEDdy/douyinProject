package controller

import (
	"douyinProject/common"
	"douyinProject/model"
	"douyinProject/service"
	"douyinProject/utils"
	"github.com/gin-gonic/gin"
	"log"
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
		CreateTime: utils.GetCurrentTimeMMDD(),
	}
	if action != "" {
		_, err := service.PublishMessage(message)
		if err != nil {
			log.Println(err.Error())
			c.JSON(500, common.Response{-1, "发送消息失败"})
			return
		}
		c.JSON(200, common.Response{0, "发送消息成功"})
		return
	}
}
func GetChatList(c *gin.Context) {
	to_user_id := c.Query("to_user_id")
	touserid, _ := strconv.ParseInt(to_user_id, 10, 64)
	message, err := service.GetChatList(touserid)
	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {
		if err != nil {
			log.Println(err.Error())
			c.JSON(500, common.MessageListResponse{
				Response: common.Response{-1, "获取聊天记录失败"},
				Messages: message,
			})
			continue
		}
		c.JSON(200, common.MessageListResponse{
			Response: common.Response{0, "获取聊天记录"},
			Messages: message,
		})
		log.Println(message)
		continue

	}

}
