/*
   @Author Ted
   @Since 2023/8/2 9:18
*/

package repo

import (
	"douyinProject/common"
	"douyinProject/model"
	"log"
)

// 发表评论：视频id，用户id，评论内容；插入comment表里
func PublishComment(user_id, video_id int64, content, time string) (model.Comment, error) {
	user, err := GetUserById(user_id)
	if err != nil {
		log.Println(err.Error())
	}
	comment := model.Comment{
		UserId:  user_id,
		VideoId: video_id,
		User:    user,
		Content: content,
		Time:    time,
	}
	//向数据库插入数据
	if err := common.DB.Create(&comment).Error; err != nil {
		log.Println(err.Error())
		return comment, err
	}
	return comment, nil
}

func DeleteComment(comment_id int64) error {
	var comment model.Comment
	if err := common.DB.Where("comment_id=?", comment_id).Delete(&comment).Error; err != nil {
		log.Println(err.Error())
		return err
	}
	return nil //删除成功
}

func GetCommentList(video_id int64) ([]*model.Comment, error) {
	var commentList []*model.Comment
	if err := common.DB.Where("video_id=?", video_id).Find(&commentList).Error; err != nil {
		log.Println(err.Error())
		return commentList, err
	}
	//根据每个comment的userid查询对应user,并封装回去
	for _, item := range commentList {
		user, err := GetUserById(item.UserId)
		if err != nil {
			log.Println(err.Error())
		}
		item.User = user
	}
	return commentList, nil
}
