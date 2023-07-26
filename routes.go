/*
   @Author Ted
   @Since 2023/7/25 14:24
*/

package main

import (
	"douyinProject/middleware"
	"github.com/gin-gonic/gin"
)
import "douyinProject/controller"

func RouteInit(r *gin.Engine) *gin.Engine {
	baseRouter := r.Group("/douyin") //路由组
	baseRouter.GET("/feed", controller.Feed)
	userRouter := baseRouter.Group("/user")
	{
		userRouter.POST("/register/", controller.Register)
		userRouter.POST("/login/", controller.Login)
		userRouter.GET("/", middleware.UserMiddleware(), controller.UserInfo)
	}

	return r
}
