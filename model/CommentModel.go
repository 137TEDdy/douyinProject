/*
   @Author Ted
   @Since 2023/7/25
*/

package model

type Comment struct {
	Id      int64  `json:"id" gorm:"column:comment_id; primary_key;"`
	UserId  int64  `json:"-"`
	User    User   `json:"user" gorm:"foreignkey:UserId"`
	VideoId int64  `json:"-"`
	Content string `json:"content"` //内容
	Time    string `json:"create_data"`
}

func (Comment) TableName() string {
	return "comments"
}
