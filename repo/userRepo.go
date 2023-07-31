/*
   @Author Ted
   @Since 2023/7/25 20:57
*/

package repo

import (
	"douyinProject/common"
	"log"
)
import "douyinProject/model"

//func IsUsernameExsit(username string) bool {
//	var user model.User
//	common.DB.Where("user_name=?", username).First(&user)
//	//nil用于比较指针类型\切片等的变量，结构体、基本类型不行！
//	if user.Id != 0 { //默认为0值，即不存在时
//		return true
//	}
//	return false
//}

// 返回user对象和bool；兼具判断user是否存在，和获取user
func GetUserById(userId int64) (model.User, error) {
	var user model.User
	if err := common.DB.First(&user, userId).Error; err != nil {
		log.Println(err.Error())
		return user, err
	}
	return user, nil
}

func GetUserByName(username string) (model.User, error) {
	var user model.User
	if err := common.DB.Where("user_name=?", username).First(&user).Error; err != nil {
		log.Println(err.Error())
		return user, err
	}
	//nil用于比较指针类型\切片等的变量，结构体、基本类型不行
	return user, nil
}

// 获取最后一位的userId，+1后用于新用户
func GetLastUserId() (int64, error) {
	var user model.User
	if err := common.DB.Last(&user).Error; err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return user.Id, nil
}

func CreateUser(username, hasedPassword string) error {
	//4.创建用户
	user := model.User{
		Name:     username,
		Password: hasedPassword,
	}
	if err := common.DB.Create(&user).Error; err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
