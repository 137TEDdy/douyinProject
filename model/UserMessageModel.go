package model

type Message struct {
	Id         int64  // 消息id
	UserId     int64  //用户id
	ToUserId   int64  //对方用户id
	ActionType int    //1-发送消息
	Content    string //消息内容
	CreateTime string // 消息创建时间
}

func (table *Message) TableName() string {
	return "Message"
}
