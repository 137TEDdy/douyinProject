/*
   @Author Ted
   @Since 2023/7/31 19:53
*/

package repo

import (
	"douyinProject/common"
	"douyinProject/model"
	"errors"
	"gorm.io/gorm"
	"log"
)

// 点赞
func Like(video_id, user_id int64) error {
	//首先为favorite填充值，否则create时这两个id值为0
	favorite := model.Favorite{
		VideoId: video_id,
		UserId:  user_id,
	}

	var video model.Video
	var err error
	err = common.DB.Where("user_id = ? and video_id = ?", user_id, video_id).Find(&favorite).Error
	//这里由于没查到，肯定有err； 只是不能是ErrRecordNotFound; （这里必须判断nil的情况，不然可能空指针错误）
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err.Error())
		return err
	}
	err = common.DB.Where("video_id = ?", video_id).Find(&video).Error
	if err != nil {
		log.Println(err.Error())
		return err
	}
	if favorite.Id != 0 {
		log.Println("已经存在该点赞记录，无法再次点赞")
		return errors.New("已经存在该点赞记录，无法再次点赞")
	}

	//这里可以换成原子操作
	video.FavoriteCount = video.FavoriteCount + 1 //点赞加1
	err = common.DB.Save(video).Error             //更新回去
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = common.DB.Create(&favorite).Error //插入记录
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func UnLike(video_id, user_id int64) error {
	var favorite model.Favorite
	var video model.Video
	var err error
	//尝试删除
	err = common.DB.Where("user_id = ? and video_id = ?", user_id, video_id).Delete(&favorite).Error
	//如果没查询到这条记录，则错误
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = common.DB.Where("video_id = ?", video_id).Find(&video).Error
	if err != nil {
		log.Println(err.Error())
		return err
	}
	if favorite.Id == 0 {
		log.Println("不存在该点赞记录，无法取消点赞")
		return errors.New("不存在该点赞记录，无法取消点赞")
	}

	video.FavoriteCount = video.FavoriteCount - 1 //点赞加1

	err = common.DB.Save(&video).Error //保存更改
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

// 查询某用户的在favorite表里所有记录
func GetFavoritesByUserid(user_id int64) ([]*model.Favorite, error) {
	var favorites []*model.Favorite

	err := common.DB.Where("user_id = ?", user_id).Find(&favorites).Error
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return favorites, nil
}

// 通过视频id,用户id，查询有没有点赞记录
func IsFavoriteExist(user_id, video_id int64) (bool, error) {
	var favorite model.Favorite
	//log.Println("")
	//尝试删除
	err := common.DB.Where("user_id = ? and video_id = ?", user_id, video_id).Find(&favorite).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err.Error())
		return false, err
	}
	if favorite.Id != 0 {
		return true, nil
	}
	return false, nil
}
