/*
   @Author Ted
   @Since 2023/7/31 15:46
*/

package controller

import (
	. "douyinProject/common"
	"douyinProject/log"
	"douyinProject/model"
	"douyinProject/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 点赞功能
func FavoriteLike(c *gin.Context) {
	//判断用户token合法性在中间件完成
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")
	usertmp, isExist := c.Get("user")
	if isExist == false {
		log.Error("根据token获取user出错")
		Resp(c, CodeTokenError, Response{-1, Msg(CodeTokenError)})
		return
	}
	user := usertmp.(model.User)

	//根据视频id，获取用户id，查favorite表；不存在就+1并插入数据，存在则删除数据并减一
	//调用service方法，repo层实现一个根据id，修改指定FavoriteCount的方法
	actionType, err := strconv.Atoi(action_type)       //转成数字
	videoId, err := strconv.ParseInt(video_id, 10, 64) //转成int64，  参数含义：十进制的64位
	if err != nil {
		log.Error(err.Error())
		Resp(c, CodeInvalidParams, Response{-1, Msg(CodeInvalidParams)})
		return
	}

	err = service.FavoriteLike(videoId, user.Id, actionType)
	if err != nil {
		log.Error(err.Error())
		Resp(c, CodeFavoriteError, Response{-1, Msg(CodeFavoriteError)})
		return
	}
	Resp(c, CodeSuccess, Response{0, Msg(CodeSuccess)})

}

// 点赞列表
func FavoriteList(c *gin.Context) {
	user_id := c.Query("user_id")
	//在favorite表里查询该用户的记录，查询出所有video_id, 依次封装到video切片里，并封装用户信息；然后返回该切片
	userId, err := strconv.ParseInt(user_id, 10, 64) //转成int64，  参数含义：十进制的64位
	if err != nil {
		log.Error(err.Error())
		Resp(c, CodeInvalidParams, Response{-1, Msg(CodeInvalidParams)})
		return
	}
	//根据userid查询视频列表
	videoList, err := service.FavoriteList(userId)
	if err != nil {
		log.Error(err.Error())
		Resp(c, CodeFavoriteError, Response{-1, Msg(CodeFavoriteError)})
		return
	}

	//响应数据
	Resp(c, CodeSuccess, VideosDto{
		Response:  Response{0, Msg(CodeSuccess)},
		VideoList: videoList,
	})

}
