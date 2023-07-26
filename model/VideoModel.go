/*
   @Author Ted
   @Since 2023/7/25
*/

package model

/*
	"id": 0,
	"author": {
		.....
	},
	"play_url": "string",  视频播放地址
	"cover_url": "string", 封面地址
	"favorite_count": 视频点赞总数
	"comment_count": 评论总数
	"is_favorite": true,   是否点赞
	"title": "string"     视频标题
*/

type Video struct {
	Id       int64 `json:"id" gorm:"column:video_id; primary_key;"`
	AuthorId int64 `json:"-"`
	Author   User  `json:"author" gorm:"foreignkey:AuthorId"`

	PlayUrl       string `json:"play_url" gorm:"column:play_url;"`
	CoverUrl      string `json:"cover_url"gorm:"column:cover_url;"`
	FavoriteCount int64  `json:"favorite_count" gorm:"column:favorite_count;"`
	CommentCount  int64  `json:"comment_count" gorm:"column:comment_count;"`
	PublishTime   int64  `gorm:"column:publish_time;" json:"-"` //不是日期类，简化成int
	IsFavorite    bool   `json:"is_favorite" gorm:"is_favorite"`
	Title         string `json:"title"`
}

// gorm 将结构体名称转换为复数形式作为表名,这里需要指定表名为users
func (Video) TableName() string {
	return "videos"
}
