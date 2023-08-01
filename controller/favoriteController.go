/*
   @Author Ted
   @Since 2023/7/31 15:46
*/

package controller

import (
	"douyinProject/common"
	"douyinProject/model"
	"douyinProject/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// 点赞功能
func FavoriteLike(c *gin.Context) {
	//判断用户token合法性在中间件完成
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")
	usertmp, isExist := c.Get("user")
	if isExist == false {
		log.Println("根据token获取user出错")
		c.JSON(500, common.Response{-1, "根据token获取user出错"})
		return
	}
	user := usertmp.(model.User)

	//根据视频id，获取用户id，查favorite表；不存在就+1并插入数据，存在则删除数据并减一
	//调用service方法，repo层实现一个根据id，修改指定FavoriteCount的方法
	actionType, err := strconv.Atoi(action_type)       //转成数字
	videoId, err := strconv.ParseInt(video_id, 10, 64) //转成int64，  参数含义：十进制的64位
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, common.Response{-1, "数字转换出错"})
		return
	}

	err = service.FavoriteLike(videoId, user.Id, actionType)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, common.Response{-1, "点赞操作出错"})
		return
	}
	c.JSON(200, common.Response{0, "操作成功"})

}
