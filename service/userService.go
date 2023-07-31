/*
   @Author Ted
   @Since 2023/7/26 9:05
*/

package service

import (
	"douyinProject/model"
	"douyinProject/repo"
)

// 返回user对象和bool；兼具判断user是否存在，和获取user
func GetUserById(userId int64) (model.User, error) {
	user, err := repo.GetUserById(userId)
	return user, err
}

func GetUserByName(username string) (model.User, error) {
	user, err := repo.GetUserByName(username)
	return user, err
}

// 获取最后一个用户的id
func GetLastUserId() (int64, error) {
	id, err := repo.GetLastUserId()
	return id, err
}

func CreateUser(username, hasedPassword string) error {
	//4.创建用户
	err := repo.CreateUser(username, hasedPassword)
	return err
}
