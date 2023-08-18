/*
   @Author Ted
   @Since 2023/7/29 20:46
*/

package repo

import (
	"douyinProject/common"
	"douyinProject/log"
	"douyinProject/model"
)

// 用户登录状态下，获取视频列表
func GetVideoListLogin(user_id int64) ([]*model.Video, error) {

	var videoList []*model.Video
	if err := common.DB.Find(&videoList).Error; err != nil {
		log.Error(err.Error())
		return videoList, err
	}
	for _, item := range videoList {
		id := item.AuthorId
		if id == 0 { //id不能为0
			continue
		}

		user, err := CacheGetUser(id) //先查找缓存
		if err != nil {               //如果缓存不存在或出错，则从数据库查找
			user, err = GetUserById(id)
			if err != nil {
				log.Error(err.Error())
			}
		}
		item.Author = user

		//判断当前用户有没有点赞，并填充is_favorite字段
		flag, err := IsFavoriteExist(user_id, item.Id)
		if err != nil {
			log.Error(err.Error())
		}
		item.IsFavorite = flag
	}
	return videoList, nil
}

// 用户没有登录状态下，获取视频列表
func GetVideoListUnLogin() ([]*model.Video, error) {

	var videoList []*model.Video
	if err := common.DB.Find(&videoList).Error; err != nil {
		log.Error(err.Error())
		return videoList, err
	}
	for _, item := range videoList {
		id := item.AuthorId
		if id == 0 { //id不能为0
			continue
		}

		user, err := CacheGetUser(id) //先查找缓存
		if err != nil {               //如果缓存不存在或出错，则从数据库查找
			user, err = GetUserById(id)
			if err != nil {
				log.Error(err.Error())
			}
		}
		item.Author = user

	}
	return videoList, nil
}

// 根据用户id获取视频列表
func GetVideoListByUserID(userId int64) ([]*model.Video, error) {
	var videoList []*model.Video
	if err := common.DB.Where("author_id=?", userId).Find(&videoList).Error; err != nil { //不要忘加&
		log.Error(err.Error())
		return videoList, err
	}
	//查询用户
	user, err := CacheGetUser(userId) //先查找缓存
	if err != nil {                   //如果缓存不存在或出错，则从数据库查找
		user, err = GetUserById(userId)
		if err != nil {
			log.Error(err.Error())
		}
	}

	for _, item := range videoList {
		id := item.AuthorId
		if id == 0 { //id不能为0
			continue
		}
		//查找该用户
		item.Author = user

		//判断当前用户有没有点赞，并填充is_favorite字段
		flag, err := IsFavoriteExist(user.Id, item.Id)
		if err != nil {
			log.Error(err.Error())
		}
		item.IsFavorite = flag
	}
	return videoList, nil
}

func StoreVideo(video model.Video) error {
	////每次保存默认为false
	//video.IsFavorite = false

	if err := common.DB.Create(&video).Error; err != nil {
		log.Error(err.Error())
		return err
	}
	//更新用户的作品数
	log.Info("作品数+1")
	UpdateUser(video.Author.Id, 1, "work_count")

	return nil
}

func UpdateVideo(videoId, num int64, stype string) error {
	video, err := GetVideosByVideoId(videoId)
	if err != nil {
		return err
	}
	switch stype {
	case "favorite_count":
		common.DB.Model(&video).Update("favorite_count", video.FavoriteCount+num)
	case "comment_count":
		common.DB.Model(&video).Update("comment_count", video.CommentCount+num)
	}
	return nil
}

// 这个方法是点赞逻辑展示核心，貌似刷视频时请求都是这个地址；
// 根据视频id获取某个视频
// 参数是当前视频id；
func GetVideosByVideoId(video_id int64) (*model.Video, error) {
	var video *model.Video
	if err := common.DB.Find(&video, video_id).Error; err != nil {
		log.Error(err.Error())
		return nil, err
	}
	//查询user相关信息并封装到video里面

	user, err := GetUserById(video.AuthorId)
	if err != nil {
		log.Error(err.Error())
	}
	video.Author = user

	return video, nil
}

func GetAuthorIdByVideoId(VideoId int64) (int64, error) {

	var video model.Video

	if err := common.DB.First(&video, VideoId).Error; err != nil {
		log.Error(err.Error())
		return 0, err
	}
	return video.AuthorId, nil
}

//// 通过视频的地址来缓存
//func CacheSetVideo(video model.Video) error {
//	id := video.Id
//	strId := strconv.Itoa(int(id))
//	err := common.CacheSet("video_"+strId, video)
//
//	return err
//}
//
//func CacheGetVideo(video_id int64) (model.Video, error) {
//	key := strconv.FormatInt(video_id, 10)
//	data, err := common.CacheGet("user_" + key)
//	var video model.Video
//	if err != nil {
//		return video, err
//	}
//
//	err = json.Unmarshal(data, &video)
//	return video, err
//}
