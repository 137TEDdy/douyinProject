/*
   @Author Ted
   @Since 2023/7/25
*/

package model

type Comment struct {
	Id      int64 `gorm:"column:comment_id; primary_key;"`
	UserId  int64
	VideoId int64
	Content string //内容
	Time    string
}

func (Comment) TableName() string {
	return "comments"
}
