package controller

import (
	. "douyinProject/common"
	"douyinProject/log"
	"douyinProject/model"
	"douyinProject/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FriendDto struct {
	Response
	FriendList []*model.User `json:"user_list,omitempty"`
}

// 关注功能
func FollowIdol(c *gin.Context) {
	//判断用户token合法性在中间件完成
	to_user_id := c.Query("to_user_id")
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
	actionType, err := strconv.Atoi(action_type)        //转成数字
	idolId, err := strconv.ParseInt(to_user_id, 10, 64) //转成int64，  参数含义：十进制的64位
	if err != nil {
		log.Error(err.Error())
		Resp(c, CodeInvalidParams, Response{-1, Msg(CodeInvalidParams)})
		return
	}
	//用户不能关注自己
	if user.Id == idolId {
		Resp(c, CodeInvalidParams, Response{
			StatusCode: -1,
			StatusMsg:  "",
		})
		return
	}

	err = service.FollowIdol(user.Id, idolId, actionType)
	if err != nil {
		log.Error(err.Error())
		Resp(c, CodeFavoriteError, Response{-1, Msg(CodeFollowError)})
		return
	}
	Resp(c, CodeSuccess, Response{0, Msg(CodeSuccess)})

}

// 获取关注列表  我的关注
func FollowList(c *gin.Context) {
	user_id := c.Query("user_id")
	//在favorite表里查询该用户的记录，查询出所有video_id, 依次封装到video切片里，并封装用户信息；然后返回该切片
	userId, err := strconv.ParseInt(user_id, 10, 64) //转成int64，  参数含义：十进制的64位
	if err != nil {
		log.Error(err.Error())
		Resp(c, CodeInvalidParams, Response{-1, Msg(CodeInvalidParams)})
		return
	}
	//根据userid查询idol列表
	idolsList, err := service.FollowsList(userId)
	if err != nil {
		log.Error(err.Error())
		Resp(c, CodeFavoriteError, Response{-1, Msg(CodeFollowError)})
		return
	}

	//响应数据
	Resp(c, CodeSuccess, UsersDto{
		Response:  Response{0, Msg(CodeSuccess)},
		UsersList: idolsList,
	})
}

// 获取被关注列表  我的粉丝
func FollowerList(c *gin.Context) {
	user_id := c.Query("user_id")
	//在favorite表里查询该用户的记录，查询出所有video_id, 依次封装到video切片里，并封装用户信息；然后返回该切片
	userId, err := strconv.ParseInt(user_id, 10, 64) //转成int64，  参数含义：十进制的64位
	if err != nil {
		log.Error(err.Error())
		Resp(c, CodeInvalidParams, Response{-1, Msg(CodeInvalidParams)})
		return
	}
	//根据userid查询idol列表
	idolsList, err := service.FollowersList(userId)
	if err != nil {
		log.Error(err.Error())
		Resp(c, CodeFavoriteError, Response{-1, Msg(CodeFollowError)})
		return
	}

	//响应数据
	Resp(c, CodeSuccess, UsersDto{
		Response:  Response{0, Msg(CodeSuccess)},
		UsersList: idolsList,
	})
}

// 获取所有关注登录用户的粉丝列表
func FriendList(c *gin.Context) {
	token := c.Query("token")
	user_id := c.Query("user_id")
	if token == "" {
		Resp(c, 500, Response{-1, "无用户token信息"})
		return
	}

	//查数据库，根据userId查出好友列表
	id, _ := strconv.ParseInt(user_id, 10, 64) //转成int64，  参数含义：十进制的64位
	userList, err := service.FriendList(id)
	if err != nil {
		log.Error(err.Error())
		Resp(c, 500, Response{0, "获取用户好友列表失败"})
		return
	}
	//log.Println("好友列表：", friendList)
	Resp(c, 200, FriendDto{
		Response:   Response{0, "获取该用户好友列表成功"},
		FriendList: userList,
	})
}
