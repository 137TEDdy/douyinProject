package service

import (
	"douyinProject/log"
	"douyinProject/model"
	"douyinProject/repo"
)

func FollowIdol(user_id, idol_id int64, action_type int) error {
	if action_type == 1 {
		err := repo.Follow(user_id, idol_id)

		if err != nil {
			return err
		}
	} else {
		err := repo.UnFollow(user_id, idol_id)
		if err != nil {
			return err
		}
	}
	return nil
}

func FollowsList(user_id int64) ([]*model.User, error) {
	//查询出该用户idol列表
	idols, err := repo.GetFollowsByUserid(user_id)
	if err != nil {
		return nil, err
	}
	//查询用户相关信息
	var userList []*model.User
	for _, val := range idols {
		user_id := val.Id
		user, err := repo.GetUserById(user_id) //获取
		if err != nil {                        //出错则不添加这条用户信息
			log.Error(err.Error())
			continue
		}
		userList = append(userList, &user)
	}
	return userList, nil
}
func FollowersList(user_id int64) ([]*model.User, error) {
	//查询出该用户粉丝列表
	idols, err := repo.GetFollowersByUserid(user_id)
	if err != nil {
		return nil, err
	}
	//查询用户相关信息
	var userList []*model.User
	for _, val := range idols {
		user_id := val.Id
		user, err := repo.GetUserById(user_id) //获取
		if err != nil {                        //出错则不添加这条用户信息
			log.Error(err.Error())
			continue
		}
		userList = append(userList, &user)
	}
	return userList, nil
}
