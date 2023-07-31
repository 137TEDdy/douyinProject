/*
   @Author Ted
   @Since 2023/7/27 10:21
*/

package utils

import "time"

// 返回当前时间
func GetCurrentTime() int64 {
	return time.Now().Unix()
}
