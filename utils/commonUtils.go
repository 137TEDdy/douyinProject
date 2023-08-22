/*
   @Author Ted
   @Since 2023/7/27 10:21
*/

package utils

import (
	"math/rand"
	"time"
)

// 返回当前时间
func GetCurrentTime() int64 {
	return time.Now().Unix()
}

// 以mm-dd格式获取当前时间
func GetCurrentTimeMMDD() string {
	curtime := time.Now()
	timeStr := curtime.Format("01-02") //格式化字符串，其中01表示月份，02表示日期
	return timeStr
}

func GetCurrentTimeForString() string {
	currentTime := time.Now()
	return currentTime.Format("200601021504")
}

func GetCurrentTimeYyyy_MM_dd_HH_MM_ss() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05")
}

func GetRandomNumber() int {
	rand.Seed(time.Now().Unix())
	// 生成随机数
	randomNum := rand.Intn(3) // 生成0~2之间的随机数
	return randomNum
}
