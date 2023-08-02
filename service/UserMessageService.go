package service

import (
	"douyinProject/model"
	"douyinProject/repo"
)

func PublishMessage(user_id int64, to_user_id int64, content string, t string) (model.Message, error) {
	message, err := repo.Publishmessage(user_id, to_user_id, content, t)
	return message, err
}
func GetChatList(to_user_id int64) ([]*model.Message, error) {
	message, err := repo.GetChatList(to_user_id)
	return message, err
}
