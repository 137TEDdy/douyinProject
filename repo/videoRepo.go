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

// 用户登录状态下，获取视频列表
func GetVideoListLogin(user_id int64) ([]*model.Video, error) {

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

		//判断当前用户有没有点赞，并填充is_favorite字段
		flag, err := IsFavoriteExist(user_id, item.Id)
		if err != nil {
			log.Println(err.Error())
		}
		item.IsFavorite = flag
	}
	return videoList, nil
}

// 用户没有登录状态下，获取视频列表
func GetVideoListUnLogin() ([]*model.Video, error) {

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

// 根据用户id获取视频列表
func GetVideoListByUserID(userId int64) ([]*model.Video, error) {
	var videoList []*model.Video
	if err := common.DB.Where("author_id=?", userId).Find(&videoList).Error; err != nil { //不要忘加&
		log.Println(err.Error())
		return videoList, err
	}
	//查询用户
	user, err := GetUserById(userId)
	if err != nil {
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

		//判断当前用户有没有点赞，并填充is_favorite字段
		flag, err := IsFavoriteExist(user.Id, item.Id)
		if err != nil {
			log.Println(err.Error())
		}
		item.IsFavorite = flag
	}
	return videoList, nil
}

func StoreVideo(video model.Video) error {
	////每次保存默认为false
	//video.IsFavorite = false

	if err := common.DB.Create(&video).Error; err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

// 这个方法是点赞逻辑展示核心，貌似刷视频时请求都是这个地址；
// 根据视频id获取某个视频
// 参数是当前视频id、以及当前登录用户id；
func GetVideosByVideoId(video_id, user_id int64) (*model.Video, error) {
	var video *model.Video
	if err := common.DB.Find(&video, video_id).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}
	//查询user相关信息并封装到video里面
	user, err := GetUserById(user_id)
	if err != nil {
		log.Println(err.Error())
	}
	video.Author = user

	//判断当前用户有没有点赞，并填充is_favorite字段
	flag, err := IsFavoriteExist(user_id, video_id)
	if err != nil {
		log.Println(err.Error())
	}
	video.IsFavorite = flag

	return video, nil
}
