package repo

import (
	"douyinProject/common"
	"douyinProject/log"
	"douyinProject/model"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

//关注或取消关注
func FollowOrUnFollowAction(follow_id int64, follower_id int64, action_type int) error{
	//isFollow为1关注，0取消关注
	//首先为relation填充值，否则create时这两个id值为0
	relation := model.Relation{
		FollowId:   follow_id,
		FollowerId: follower_id,
	}
	var err error
	err = common.DB.Where("follow_id = ? and follower_id = ?", follow_id, follower_id).Take(&relation).Error
	if action_type==1 {
		//关注
		//这里由于没查到，肯定有err； 只是不能是ErrRecordNotFound; （这里必须判断nil的情况，不然可能空指针错误）
		if err != nil && err != gorm.ErrRecordNotFound {
			log.Error(err.Error())
			return err
		}
		if relation.Id != 0 {
			log.Error("已经存在该关注记录，无法再次关注")
			return errors.New("已经存在该关注记录，无法再次关注")
		}

		//插入记录
		err = common.DB.Create(&relation).Error
		if err != nil {
			log.Error(err.Error())
			return err
		}
	}else{

		if relation.Id != 0 {
			log.Error(follow_id,"未关注",follower_id)
			return errors.New("未关注，请先关注")
		}

		//删除记录
		err=common.DB.Delete(&relation).Error
		if err != nil {
			log.Error(err.Error())
			return err
		}
	}

	var op int64;
	if action_type==1 {
		op=add1
	}else{
		op=sub1;
	}
	//更新用户的关注和被关注信息
	UpdateUser(follow_id, op, "following_count")

	UpdateUser(follower_id, op, "followers_count")

	go CacheChangeUserCount(follow_id, op, "follow")
	go CacheChangeUserCount(follower_id, op, "followed")
	return nil
}



// 查询某用户的在favorite表里所有记录 查看我的关注
func GetFollowsByUserid(follow_id int64) ([]*model.Relation, error) {
	var relations []*model.Relation

	err := common.DB.Where("follow_id = ?", follow_id).Find(&relations).Error
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return relations, nil
}

// 查询某用户的在favorite表里所有记录 查看我的粉丝
func GetFollowersByUserid(follower_id int64) ([]*model.Relation, error) {
	var relations []*model.Relation

	err := common.DB.Where("follower_id = ?", follower_id).Find(&relations).Error
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return relations, nil
}

func CacheChangeFollowsCount(userid, op int64, ftype string) {
	uid := strconv.FormatInt(userid, 10)
	mutex, _ := common.GetLock("user_" + uid) //获取锁，最后释放
	defer common.UnLock(mutex)
	user, err := CacheGetUser(userid)
	if err != nil {
		log.Error("user:%v miss cache", userid)
		return
	}
	switch ftype {
	case "follow":
		user.FollowCount += op
	case "followed":
		user.FollowerCount += op
	}
	CacheSetUser(user)
}
