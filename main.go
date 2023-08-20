package main

import (
	"douyinProject/common"
	"douyinProject/config"
	"douyinProject/log"
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

	err := r.Run(":9093")
	if err != nil {
		return
	}

}
