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
}

func main() {
	r := gin.Default()
	RouteInit(r)
	Init()
	r.Run(":9093")
}
