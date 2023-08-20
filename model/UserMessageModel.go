package model

type Message struct {
	Id         int64  `json:"id" gorm:"column:comment_id; primary_key;"` // 消息id
	UserId     int64  //用户id
	ToUserId   int64  //对方用户id
	ActionType int    //1-发送消息
	Content    string `json:"content"`     //消息内容
	CreateTime string `json:"create_data"` // 消息创建时间
}

func (table *Message) TableName() string {
	return "Message"
}
