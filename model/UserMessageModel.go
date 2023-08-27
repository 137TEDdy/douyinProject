package model

type Message struct {
	Id         int64  `json:"id" gorm:"column:message_id; primary_key;"` // 消息id
	UserId     int64  `json:"from_user_id"`                              //用户id
	ToUserId   int64  `json:"to_user_id"`                                //对方用户id
	ActionType int    `json:"-"`                                         //1-发送消息
	Content    string `json:"content"`                                   //消息内容
	CreateTime string `json:"create_time"`                               // 消息创建时间
}

func (table *Message) TableName() string {
	return "Message"
}
