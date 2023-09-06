package main

import (
	"douyinProject/common"
	"douyinProject/config"
	"douyinProject/controller"
	"douyinProject/log"
	"douyinProject/middleware"
	"douyinProject/minioHandler"

	"github.com/gin-gonic/gin"
)

func Init() {

	config.InitConfig() //初始化配置
	log.InitLog()
	minioHandler.InitMinio()
	common.DBInit()
	common.RedisInit()

}

func main() {
	r := gin.Default()
	RouteInit(r)
	Init()

	err := r.Run(":4000")
	if err != nil {
		log.Error("启动错误")
		return
	}

}

func RouteInit(r *gin.Engine) *gin.Engine {
	baseRouter := r.Group("/douyin") //路由组
	baseRouter.GET("/feed", controller.Feed)
	commentRouter := baseRouter.Group("/comment")
	{
		commentRouter.POST("/action/", middleware.TokenMiddleware(), controller.CommentAction)
		commentRouter.GET("/list/", middleware.TokenMiddleware(), controller.GetCommentList)
	}
	favoriteRouter := baseRouter.Group("/favorite")
	{
		favoriteRouter.POST("/action/", middleware.TokenMiddleware(), controller.FavoriteLike)
		favoriteRouter.GET("/list/", middleware.TokenMiddleware(), controller.FavoriteList)
	}
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

	relationRouter := baseRouter.Group("/relation")
	{
		relationRouter.POST("/action/", middleware.TokenMiddleware(), controller.FollowIdol)
		relationRouter.GET("/follow/list/", middleware.TokenMiddleware(), controller.FollowList)
		relationRouter.GET("/follower/list/", middleware.TokenMiddleware(), controller.FollowerList)
		relationRouter.GET("/friend/list/", middleware.TokenMiddleware(), controller.FriendList)
	}

	messageRouter := baseRouter.Group("/message")
	{
		messageRouter.POST("/action/", middleware.TokenMiddleware(), controller.ChatAction)
		messageRouter.GET("/chat/", middleware.TokenMiddleware(), controller.GetChatList)
	}
	return r
}
