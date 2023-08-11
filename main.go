package main

import (
	"douyinProject/common"
	"douyinProject/config"
	"douyinProject/minioHandler"
	"github.com/gin-gonic/gin"
)

func Init() {

	config.InitConfig() //初始化配置
	minioHandler.InitMinio()
	common.DBInit()
	common.RedisInit()

}

func main() {
	r := gin.Default()
	RouteInit(r)
	Init()
	r.Run(":9093")
	//common.CacheHSet("comment_1", "2", "xxx")
	//common.CacheHSet("comment_1", "3", "aaa")
	//comment, _ := repo.CacheGetComment(1)
	//log.Println(comment)
}
