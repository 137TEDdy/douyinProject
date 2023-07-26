/*
   @Author Ted
   @Since 2023/7/25 20:57
*/

package utils

import "douyinProject/common"
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
func GetUserById(userId int64) (model.User, bool) {
	var user model.User
	common.DB.First(&user, userId)
	if user.Id != 0 { //默认为0值，即不存在时
		return user, true
	}
	return user, false
}

func GetUserByName(username string) (model.User, bool) {
	var user model.User
	common.DB.Where("user_name=?", username).First(&user)
	//nil用于比较指针类型\切片等的变量，结构体、基本类型不行
	if user.Id != 0 { //默认为0值，即不存在时
		return user, true
	}
	return user, false
}

// 获取最后一位的userId，+1后用于新用户
func GetLastUserId() int64 {
	var user model.User
	common.DB.Last(&user)
	return user.Id
}
