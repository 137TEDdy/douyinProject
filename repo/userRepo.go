/*
   @Author Ted
   @Since 2023/7/25 20:57
*/

package repo

import (
	"douyinProject/common"
	"douyinProject/log"
	"douyinProject/model"
	"encoding/json"
	"strconv"
)

//func IsUsernameExsit(username string) bool {
//	var user model.User
//	common.DB.Where("user_name=?", username).First(&user)
//	//nil用于比较指针类型\切片等的变量，结构体、基本类型不行！
//	if user.Id != 0 { //默认为0值，即不存在时
//		return true
//	}
//	return false
//}

// 更新的用户，num为更新值
func UpdateUser(userId, num int64, stype string) {

	user, _ := GetUserById(userId)
	switch stype {
	case "favorite_count":
		common.DB.Model(&user).Update("favorite_count", user.FavCount+num)
	case "total_favorited":
		common.DB.Model(&user).Update("total_favorited", user.TotalFav+num)
	case "work_count":
		common.DB.Model(&user).Update("work_count", user.WorkCount+num)
	case "following_count":
		common.DB.Model(&user).Update("following_count", user.FollowCount+num)
	case "followers_count":
		common.DB.Model(&user).Update("followers_count", user.FollowerCount+num)
	}
}

// 返回user对象和bool；兼具判断user是否存在，和获取user
func GetUserById(userId int64) (model.User, error) {
	//var user model.User
	user, err := CacheGetUser(userId)
	//从缓存获取user成功
	if err == nil {
		return user, nil
	}
	if err := common.DB.First(&user, userId).Error; err != nil {
		log.Error(err.Error())
		return user, err
	}
	//开启协程，进行缓存
	go CacheSetUser(user)
	return user, nil
}

func GetUserByName(username string) (model.User, error) {
	var user model.User
	if err := common.DB.Where("user_name=?", username).First(&user).Error; err != nil {
		log.Error(err.Error())
		return user, err
	}
	//开启协程，进行缓存
	go CacheSetUser(user)
	//nil用于比较指针类型\切片等的变量，结构体、基本类型不行
	return user, nil
}

// 获取最后一位的userId，+1后用于新用户
func GetLastUserId() (int64, error) {
	var user model.User
	if err := common.DB.Last(&user).Error; err != nil {
		log.Error(err.Error())
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
		log.Error(err.Error())
		return err
	}
	//开启协程，进行缓存
	go CacheSetUser(user)
	return nil
}

// 返回值为user的缓存方法封装
func CacheSetUser(user model.User) {
	id := user.Id
	strId := strconv.Itoa(int(id))              //先把int64转成int,再转成字符串
	err := common.CacheSet("user_"+strId, user) //会先序列化再cache
	if err != nil {
		log.Error("缓存失败，", err.Error())
	}
}

func CacheGetUser(uid int64) (model.User, error) {
	key := strconv.FormatInt(uid, 10) //将int64类型的整数转换为对应的字符串表示
	data, err := common.CacheGet("user_" + key)
	var user model.User
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(data, &user) //反序列化成user

	return user, err
}
