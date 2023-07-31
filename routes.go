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
	publishRouter := baseRouter.Group("/publish")
	{
		publishRouter.POST("/action/", middleware.TokenMiddleware(), controller.Publish)
		publishRouter.GET("/list/", middleware.TokenMiddleware(), controller.GetUserVideoList)
	}

	userRouter := baseRouter.Group("/user")
	{
		userRouter.POST("/register/", controller.Register)
		userRouter.POST("/login/", controller.Login)
		userRouter.GET("/", middleware.TokenMiddleware(), controller.UserInfo)
	}

	return r
}
