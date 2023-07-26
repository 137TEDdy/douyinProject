/*
   @Author Ted
   @Since 2023/7/25
*/

package model

type Favorite struct {
	Id      int64 `gorm:"column:favorite_id; primary_key;"`
	UserId  int64
	VideoId int64
}

func (Favorite) TableName() string {
	return "favorites"
}
