/*
   @Author Ted
   @Since 2023/7/25 19:30
*/

package service

import "fmt"
import . "douyinProject/model"
import . "douyinProject/common"

func GetVideoList() []*Video {
	var videoList []*Video

	//DB是gorm.DB，是common包下的database.go的全局变量，这里直接使用
	DB.Find(&videoList) //!:  可以加上limit
	for _, item := range videoList {
		id := item.AuthorId
		var user User
		if id == 0 { //id不能为0
			continue
		}
		DB.Where("user_id=?", id).Take(&user) //查找该用户
		fmt.Println("User:", user)
		item.Author = user
	}
	return videoList
}
