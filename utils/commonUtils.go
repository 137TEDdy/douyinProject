/*
   @Author Ted
   @Since 2023/7/27 10:21
*/

package utils

import "time"

func GetCurrentTime() int64 {
	return time.Now().Unix()
}
