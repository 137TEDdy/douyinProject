/*
   @Author Ted
   @Since 2023/7/25
*/

package controller

import (
	"douyinProject/common"
	. "douyinProject/model"
	"douyinProject/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

/*
	功能：注册；
	实现： 判断用户token是否存在，然后分支处理；不存在则注册并返回token
*/

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//fmt.Println(username, "  ", password)

	//token := username + password
	//2.验证： 判断username是否存在,  存在err说明没有这个用户名
	if _, err := service.GetUserByName(username); err == nil {
		c.JSON(422, common.Response{-1, "用户已经存在"})
		log.Println("用户已经存在")
		return
	}
	//对密码进行加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, common.Response{-1, "加密错误"})
		log.Println("加密错误")
		return
	}
	//fmt.Println("加密密码：", hasedPassword)
	//4.创建用户
	service.CreateUser(username, string(hasedPassword))

	user, err := service.GetUserByName(username)
	token, err := common.ReleaseToken(user) //获取随机token
	LastUserId := user.Id
	//查询最后一位用户的id，用于自增+1作为新用户的id
	if err != nil {
		c.JSON(500, common.Response{-1, "获取用户信息或token错误"})
		log.Println("获取用户信息或token错误")
		return
	}
	//fmt.Println("user：", user)
	//.返回结果
	c.JSON(200, common.UserLoginResponse{
		common.Response{0, "注册成功"},
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
		c.JSON(422, common.Response{-1, "用户不存在"})
		log.Println(err.Error())
		return
	} else {
		user = userTmp //如果存在，则为user赋值
	}

	//解密（转为byte切片），对比密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		//如果有err，说明密码错误
		c.JSON(400, common.Response{-1, "密码错误"})
		log.Println("密码错误")
		return
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.Response{-1, "token发放失败"})
		log.Printf("token生成错误: %v", err)
		return
	}
	//3.返回
	c.JSON(200, common.UserLoginResponse{
		common.Response{0, "登录成功"},
		user.Id,
		token,
	})
}

// 功能：获取用户信息
func UserInfo(c *gin.Context) {
	fmt.Println("进入userInfo")
	user, _ := c.Get("user")

	c.JSON(200, common.UserResponse{
		common.Response{
			0,
			"获取用户信息成功",
		},
		user.(User), //类型转换，要返回User，从any->User
	})
}
