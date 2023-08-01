/*
   @Author Ted
   @Since 2023/7/31 20:02
*/

package service

import "douyinProject/repo"

func FavoriteLike(video_id, user_id int64, action_type int) error {
	if action_type == 1 {
		err := repo.Like(video_id, user_id)
		if err != nil {
			return err
		}
	} else {
		err := repo.UnLike(video_id, user_id)
		if err != nil {
			return err
		}
	}
	return nil

}
