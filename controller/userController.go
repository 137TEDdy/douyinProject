/*
   @Author Ted
   @Since 2023/7/25
*/

package controller

import (
	. "douyinProject/common"
	"douyinProject/log"
	. "douyinProject/model"
	"douyinProject/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

/*
	功能：注册；
	实现： 判断用户token是否存在，然后分支处理；不存在则注册并返回token
*/

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	//2.验证： 判断username是否存在,  存在err说明没有这个用户名
	if _, err := service.GetUserByName(username); err == nil {
		Resp(c, CodeUserExist, Response{-1, Msg(CodeUserExist)})
		log.Error("用户已经存在")
		return
	}
	//对密码进行加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		Resp(c, CodeServerError, Response{-1, Msg(CodeServerError)})
		log.Error("加密错误")
		return
	}
	//4.创建用户
	service.CreateUser(username, string(hasedPassword))

	user, err := service.GetUserByName(username)
	token, err := ReleaseToken(user) //获取随机token
	LastUserId := user.Id
	//查询最后一位用户的id，用于自增+1作为新用户的id
	if err != nil {
		Resp(c, CodeTokenError, Response{-1, Msg(CodeTokenError)})
		log.Error("获取用户信息或token错误")
		return
	}
	Resp(c, CodeSuccess, UserLoginResponse{
		Response{0, Msg(CodeSuccess)},
		LastUserId,
		token,
	})
}

/*
功能： 登录
请求方式: post,
返回值： "status_code": 0,
	"status_msg": "string",
	"user_id": 0,
	"token": "string"
*/

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	//2.验证： 判断username是否存在
	var user User
	if userTmp, err := service.GetUserByName(username); err != nil {
		Resp(c, CodeUserNotExist, Response{-1, Msg(CodeUserNotExist)})
		log.Error(err.Error())
		return
	} else {
		user = userTmp //如果存在，则为user赋值
	}

	//解密（转为byte切片），对比密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		//如果有err，说明密码错误
		Resp(c, CodePasswordError, Response{-1, Msg(CodePasswordError)})
		log.Error("密码错误")
		return
	}
	//发放token
	token, err := ReleaseToken(user)
	if err != nil {
		Resp(c, CodeTokenError, Response{-1, Msg(CodeTokenError)})
		log.Error("token生成错误: %v", err)
		return
	}
	//3.返回
	//c.JSON(CodeSuccess, UserLoginResponse{
	//	Response{0, Msg(CodeSuccess)},
	//	user.Id,
	//	token,
	//})
	Resp(c, CodeSuccess, UserLoginResponse{
		Response{0, Msg(CodeSuccess)},
		user.Id,
		token,
	})
}

// 功能：获取用户信息
func UserInfo(c *gin.Context) {
	user, flag := c.Get("user")
	if flag == false {
		Resp(c, CodeTokenError, Response{-1, Msg(CodeTokenError)})
		return
	}

	Resp(c, CodeSuccess, UserResponse{
		Response{
			0,
			Msg(CodeSuccess),
		},
		user.(User), //类型转换，要返回User，从any->User
	})
}

// 封装user的关注列表请求返回值
type UsersDto struct {
	Response
	NextTime  int64   `json:"next_time,omitempty"`
	UsersList []*User `json:"user_list,omitempty"`
}
