package service

import (
	"douyinProject/model"
	"douyinProject/repo"
)

func PublishMessage(message model.Message) (bool, error) {
	ok, err := repo.Publishmessage(message)
	return ok, err
}
func GetChatList(to_user_id int64) ([]*model.Message, error) {
	message, err := repo.GetChatList(to_user_id)
	return message, err
}
