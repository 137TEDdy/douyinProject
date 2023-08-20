package service

import (
	"douyinProject/model"
	"douyinProject/repo"
)

func PublishMessage(message model.Message) (bool, error) {
	ok, err := repo.Publishmessage(message)
	return ok, err
}
func GetChatList(user_id, to_user_id int64, time string) ([]*model.Message, error) {
	message, err := repo.GetChatList(user_id, to_user_id, time)
	return message, err
}
