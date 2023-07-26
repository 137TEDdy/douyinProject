package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("我们的抖音项目")
	//common.DBInit()
	r := gin.Default()
	RouteInit(r)
	r.Run(":9093")
}
