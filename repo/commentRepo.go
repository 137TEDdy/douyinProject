/*
   @Author Ted
   @Since 2023/8/2 9:18
*/

package repo

import (
	"douyinProject/common"
	"douyinProject/log"
	"douyinProject/model"
	"encoding/json"
	"strconv"
)

// 发表评论：视频id，用户id，评论内容；插入comment表里
func PublishComment(user_id, video_id int64, content, time string) (model.Comment, error) {
	user, err := CacheGetUser(user_id) //先查找缓存
	if err != nil {                    //如果缓存不存在或出错，则从数据库查找
		user, err = GetUserById(user_id)
		if err != nil {
			log.Error(err.Error())
		}
	}

	comment := model.Comment{
		UserId:  user_id,
		VideoId: video_id,
		User:    user,
		Content: content,
		Time:    time,
	}
	//传入视频id，修改评论数
	UpdateVideo(video_id, 1, "comment_count")

	//向数据库插入数据
	//在if语句中的err只在if语句的作用域中有效
	if err := common.DB.Create(&comment).Error; err != nil {
		log.Error(err.Error())
		return comment, err
	}
	//缓存
	go CacheSetComment(video_id, comment)
	//必须填充comment的id属性，不然id为0；
	log.Info("comment的id是：", comment.Id)
	return comment, nil

}

func GetCommentByID(comment_id int64) (model.Comment, error) {
	//先查出该comment
	var comment model.Comment
	if err := common.DB.Take(&comment, comment_id).Error; err != nil {
		log.Error(err.Error())
		return comment, err
	}
	user, err := GetUserById(comment.UserId)
	if err != nil {
		return comment, err
	}
	comment.User = user
	return comment, nil
}

func DeleteComment(comment_id int64) error {

	//先查询comment
	log.Info("开始准备删除评论")
	comment, err := GetCommentByID(comment_id)
	//如果缓存存在则删除它
	if comment.Id != 0 {
		err = CacheDelComment(comment.VideoId, comment)
		if err != nil {
			log.Error("缓存删除失败")
			return err
		}
		log.Info("缓存删除成功")
	}

	if err := common.DB.Where("comment_id=?", comment_id).Delete(&comment).Error; err != nil {
		log.Error(err.Error())
		return err
	}
	//传入视频id，修改评论数
	go UpdateVideo(comment.VideoId, -1, "comment_count")
	log.Info("删除评论成功")
	return nil //删除成功
}

func GetCommentList(video_id int64) ([]*model.Comment, error) {
	var commentList []*model.Comment
	var list []*model.Comment
	//从缓存获取评论列表list
	list, err := CacheGetComment(video_id)
	if list != nil {
		log.Info("comment缓存获取成功")
		for _, val := range list {
			commentList = append(commentList, val)
		}
		return commentList, nil
	}
	if err != nil {
		log.Error("comment缓存获取出错", err.Error())
	}
	//缓存里没有，则查询mysql数据库
	if err := common.DB.Where("video_id=?", video_id).Find(&commentList).Error; err != nil {
		log.Error(err.Error())
		return commentList, err
	}

	//根据每个comment的userid查询对应user,并封装回去
	for _, item := range commentList {
		id := item.UserId
		user, err := CacheGetUser(id) //先查找缓存
		if err != nil {               //如果缓存不存在或出错，则从数据库查找
			user, err = GetUserById(id)
			if err != nil {
				log.Error(err.Error())
			}
		}

		item.User = user
		//到这里的是因为没有缓存过的，因此缓存数据
		go CacheSetComment(video_id, *item)
	}
	return commentList, nil
}

func CacheGetComment(video_id int64) ([]*model.Comment, error) {
	videoId := strconv.FormatInt(video_id, 10) //转字符串

	//以lrange的方式获取列表数据
	data, err := common.CacheLGetAll("comment_" + videoId)
	var comments []*model.Comment
	if err != nil {
		return comments, err
	}
	//反序列化
	for _, val := range data {
		var comment model.Comment
		err = json.Unmarshal(val, &comment)
		if err != nil {
			log.Error(err.Error())
		}
		comments = append(comments, &comment)
	}
	return comments, err
}

func CacheSetComment(video_id int64, comment model.Comment) error {
	videoId := strconv.FormatInt(video_id, 10)
	//以lpush的方式插入数据
	err := common.CacheLPush("comment_"+videoId, comment)
	return err
}

func CacheDelComment(video_id int64, comment model.Comment) error {
	videoId := strconv.FormatInt(video_id, 10)
	//以lrem删除数据
	// 将对象序列化为字符串
	data, err := json.Marshal(comment)
	if err != nil {
		log.Error("序列化失败:", err)
		return err
	}
	jsonString := string(data)
	err = common.CacheLRem("comment_"+videoId, jsonString)
	return err
}
