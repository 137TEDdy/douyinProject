/*
   @Author Ted
   @Since 2023/7/25
*/

package model

type Relation struct {
	Id         int64 `gorm:"column:relation_id; primary_key;"`
	FollowId   int64 //被关注者
	FollowerId int64 //关注者
}

func (Relation) TableName() string {
	return "relations"
}
