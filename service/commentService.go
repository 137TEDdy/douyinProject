/*
   @Author Ted
   @Since 2023/8/2 9:23
*/

package service

import (
	"douyinProject/model"
	"douyinProject/repo"
)

// 发表评论：视频id，用户id，评论内容；插入comment表里
func PublishComment(user_id, video_id int64, content, time string) (model.Comment, error) {
	comment, err := repo.PublishComment(user_id, video_id, content, time)
	return comment, err
}
func DeleteComment(comment_id int64) error {
	err := repo.DeleteComment(comment_id)
	return err
}
func GetCommentList(video_id int64) ([]*model.Comment, error) {
	commentlist, err := repo.GetCommentList(video_id)
	return commentlist, err

}
