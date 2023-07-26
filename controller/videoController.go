/*
   @Author Ted
   @Since 2023/7/25 14:28
*/

package controller

import (
	"douyinProject/common"
	. "douyinProject/model"
	"douyinProject/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 封装video的feed请求返回值
type VideosDto struct {
	common.Response
	VideoList []*Video `json:"video_list,omitempty"`
	NextTime  int64    `json:"next_time,omitempty"`
}

//var DemoVideos = []Video{
//	{
//		Id:            1,
//		Author:        DemoUser,
//		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
//		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
//		FavoriteCount: 0,
//		CommentCount:  0,
//		IsFavorite:    false,
//	},
//}

func Feed(c *gin.Context) {
	//time := c.Query("latest_time")
	//token:=c.Query("token")
	fmt.Println("进入feed方法")
	var videoList = service.GetVideoList() //调用service业务方法

	//fmt.Println(videoList)
	//videoList, _ := json.Marshal(videoListDemo)
	//fmt.Println(string(videoList))

	fmt.Println()
	c.JSON(200, VideosDto{
		common.Response{0, "成功"},
		videoList,
		time.Now().Unix(),
	})

}
