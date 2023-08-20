package repo

import (
	"douyinProject/common"
	"douyinProject/model"
	"fmt"
	"log"
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
		log.Println(err.Error())
		return false, err
	}
	fmt.Println(message)
	return true, nil
}
func GetChatList(to_user_id int64) ([]*model.Message, error) {
	var messageList []*model.Message
	exist, err := GetUserById(to_user_id)
	if exist == (model.User{}) || err != nil {
		return nil, err
	}
	if err = common.DB.Where("to_user_id=?", to_user_id).Find(&messageList).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return messageList, err
}
