package repo

import (
	"douyinProject/common"
	"douyinProject/log"
	"douyinProject/model"
	"fmt"
)

func Publishmessage(message model.Message) (bool, error) {
	exist, err := GetUserById(message.UserId)
	if exist == (model.User{}) || err != nil {
		return false, err
	}
	exist, err = GetUserById(message.ToUserId)
	if exist == (model.User{}) || err != nil {
		return false, err
	}
	if err := common.DB.Create(&message).Error; err != nil {
		log.Info(err.Error())
		return false, err
	}
	fmt.Println(message)
	return true, nil
}
func GetChatList(user_id, to_user_id int64, currTime string) ([]*model.Message, error) {
	var messageList []*model.Message
	exist, err := GetUserById(to_user_id)
	if exist == (model.User{}) || err != nil {
		return nil, err
	}
	//这里需要两个user_id才能唯一锁定相关记录，只返回上一条记录之后的消息，避免重复展示
	if err = common.DB.Where("((user_id=? and to_user_id=? ) or(user_id=? and to_user_id=?)) and create_time>?", user_id, to_user_id, to_user_id, user_id, currTime).Find(&messageList).Error; err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return messageList, err
}
