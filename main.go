package main

import (
	"douyinProject/common"
	"douyinProject/config"
	"douyinProject/minioHandler"
	"github.com/gin-gonic/gin"
)

func Init() {

	config.InitConfig() //初始化配置
	common.DBInit()
	minioHandler.InitMinio()

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
