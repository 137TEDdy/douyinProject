package repo

import (
	"douyinProject/common"
	"douyinProject/model"
	"log"
)

func Publishmessage(user_id int64, to_user_id int64, content string, t string) (model.Message, error) {
	message := model.Message{
		UserId:     user_id,
		ToUserId:   to_user_id,
		Content:    content,
		CreateTime: t,
	}
	if err := common.DB.Create(&message).Error; err != nil {
		log.Println(err.Error())
		return message, err
	}
	return message, nil
}
func GetChatList(to_user_id int64) ([]*model.Message, error) {
	var messageList []*model.Message
	var err error
	if err = common.DB.Where("to_user_id=?", to_user_id).Find(&messageList).Error; err != nil {
		log.Println(err.Error())
		return messageList, err
	}
	return messageList, err
}
