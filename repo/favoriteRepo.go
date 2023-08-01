/*
   @Author Ted
   @Since 2023/7/31 19:53
*/

package repo

import (
	"douyinProject/common"
	"douyinProject/model"
	"gorm.io/gorm"
)

// 点赞
func Like(video_id, user_id int64) error {
	var favorite model.Favorite
	var video model.Video
	var err error
	err = common.DB.Where("user_id = ? and video_id = ?", user_id, video_id).Find(&favorite).Error
	//这里由于没查到，肯定有err； 只是不能是ErrRecordNotFound;
	if err != gorm.ErrRecordNotFound {
		return err
	}
	err = common.DB.Where("video_id = ?", video_id).Find(&video).Error
	if err != nil {
		return err
	}

	video.FavoriteCount = video.FavoriteCount + 1 //点赞加1
	err = common.DB.Save(video).Error             //更新回去
	if err != nil {
		return err
	}

	err = common.DB.Create(&favorite).Error //插入记录
	if err != nil {
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
		return err
	}

	err = common.DB.Where("video_id = ?", video_id).Find(&video).Error
	if err != nil {
		return err
	}
	video.FavoriteCount = video.FavoriteCount - 1 //点赞加1

	err = common.DB.Save(&video).Error //保存更改
	if err != nil {
		return err
	}
	return nil

}
