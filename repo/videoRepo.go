/*
   @Author Ted
   @Since 2023/7/29 20:46
*/

package repo

import (
	"douyinProject/common"
	"douyinProject/model"
	"fmt"
	"log"
)

func GetVideoList() ([]*model.Video, error) {
	var videoList []*model.Video
	if err := common.DB.Find(&videoList).Error; err != nil {
		log.Println(err.Error())
		return videoList, err
	}
	for _, item := range videoList {
		id := item.AuthorId
		if id == 0 { //id不能为0
			continue
		}
		user, err := GetUserById(id)
		if err != nil { //用户不存在，不结束请求
			log.Println(err.Error())
			continue
		}
		item.Author = user
	}
	return videoList, nil
}

func GetVideoListByUserID(userId int) ([]*model.Video, error) {
	var videoList []*model.Video
	if err := common.DB.Where("author_id=?", userId).Find(&videoList).Error; err != nil { //不要忘加&
		log.Println(err.Error())
		return videoList, err
	}
	var user model.User
	if err := common.DB.Where("user_id=?", userId).Take(&user).Error; err != nil {
		log.Println(err.Error())
	}
	for _, item := range videoList {
		id := item.AuthorId
		if id == 0 { //id不能为0
			continue
		}
		//查找该用户
		fmt.Println("User:", user)
		item.Author = user
	}
	return videoList, nil
}

func StoreVideo(video model.Video) error {
	if err := common.DB.Create(&video).Error; err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
