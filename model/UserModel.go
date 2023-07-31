/*
   @Author Ted
   @Since 2023/7/25
*/

package model

/*
	请求返回
	"user": {
		"id": integer, 用户id
		"name": "string",  名称
		"follow_count":   关注总数
		"follower_count":   粉丝总数
		"avatar": "string",   用户头像
		"background_image": "string",   封面地址
		"signature": "string",    个人签名
		"total_favorited": "string",   获赞总数
		"work_count":       总作品数
		"favorite_count":    喜欢数
	}
*/

type User struct {
	// gorm.Model
	Id              int64  `json:"id" gorm:"column:user_id; primaryKey"`
	Name            string `json:"name" gorm:"column:user_name"`
	Password        string `json:"-"` //密码在序列化时忽略
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFav        int64  `gorm:"column:total_favorited" json:"total_favorited"`
	FavCount        int64  `gorm:"column:favorite_count" json:"favorite_count"`
	WorkCount       int64  `json:"work_count"`
	IsFollow        bool   `json:"is_follow"`
}

// gorm 将结构体名称转换为复数形式作为表名,这里需要指定表名为users
func (User) TableName() string {
	return "users"
}
