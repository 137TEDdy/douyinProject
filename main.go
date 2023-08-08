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

	//log.Println(common.Exists("name"))
	//common.CacheSet("user", "tedd")
	//config.InitConfig() //初始化配置
	//common.RedisInit()
	//nums := []int{1, 2, 3, 8, 9}
	//common.CacheHSet("comment", "1", nums)

	//common.CacheHGet("comment", "1")
}
