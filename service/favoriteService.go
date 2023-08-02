/*
   @Author Ted
   @Since 2023/7/31 20:02
*/

package service

import (
	"douyinProject/model"
	"douyinProject/repo"
	"log"
)

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

func FavoriteList(user_id int64) ([]*model.Video, error) {
	//查询出该用户点赞的所有视频的favorite列表
	favorites, err := repo.GetFavoritesByUserid(user_id)
	if err != nil {
		return nil, err
	}
	//查询视频id相关信息
	var videoList []*model.Video
	for _, val := range favorites {
		video_id := val.VideoId
		video, err := repo.GetVideosByVideoId(video_id, user_id) //获取
		if err != nil {                                          //出错则不添加这条视频信息
			log.Println(err.Error())
			continue
		}
		videoList = append(videoList, video)
	}
	return videoList, nil

}
