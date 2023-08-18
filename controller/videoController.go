/*
   @Author Ted
   @Since 2023/7/25 14:28
*/

package controller

import (
	. "douyinProject/common"
	"douyinProject/config"
	"douyinProject/log"
	"douyinProject/minioHandler"
	. "douyinProject/model"
	"douyinProject/service"
	"douyinProject/utils"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strconv"
)

// 封装video的feed请求返回值
type VideosDto struct {
	Response
	NextTime  int64    `json:"next_time,omitempty"`
	VideoList []*Video `json:"video_list,omitempty"`
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
	//判断用户是否登陆状态, 是则user_id不为0
	var user User //当前登录的用户
	var videoList []*Video
	var err error
	userTmp, _ := c.Get("user")
	//根据用户是否登录来分段处理
	//未登录，这里不是“”，而是nil
	if userTmp == nil {
		videoList, err = service.GetVideoList(0) //调用service业务方法
	} else { //已经登录
		user = userTmp.(User)
		videoList, err = service.GetVideoList(user.Id) //调用service业务方法
	}

	if err != nil {
		log.Error(err.Error())
		c.JSON(CodeGetVideoListError,
			Response{-1, Msg(CodeGetVideoListError)})
	}

	c.JSON(CodeSuccess, VideosDto{
		Response{0, Msg(CodeSuccess)},
		utils.GetCurrentTime(),
		videoList,
	})

}

// 请求路径：POST，  /douyin/publish/action/
// UploadFile(filetype, filePath, userID string) (string, error)
// 先获取文件，存到本地，再上传
func Publish(c *gin.Context) {

	user, _ := c.Get("user")
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	user_id := user.(User).Id

	filename := filepath.Base(data.Filename)
	videoBasePath := config.GetConfig().Path.VideoBasePath
	fileFinalPath := filepath.Join(videoBasePath, filename) //文件最终保存在本地的位置
	if err != nil {
		log.Error("获取文件失败: ")
		c.JSON(CodeServerError, Response{-1, Msg(CodeServerError)})
		return
	}

	//保存文件
	if err := c.SaveUploadedFile(data, fileFinalPath); err != nil {
		log.Error("保存本地失败 ")
		c.JSON(CodeServerError, Response{-1, Msg(CodeServerError)})
		return
	}
	log.Info("视频已经保存到本地")

	client := minioHandler.GetClient()
	videoUrl, err := client.UploadFile(user_id, "video", fileFinalPath)
	if err != nil {
		log.Error("上传视频失败")
		c.JSON(CodeServerError, Response{-1, Msg(CodeServerError)})
		return
	}

	//获取封面, 并上传:
	coverUrlTmp := service.GetImage(fileFinalPath)

	coverUrl, err := client.UploadFile(user_id, "image", coverUrlTmp)
	if err != nil {
		log.Error("上传图片背景失败")
		c.JSON(CodeServerError, Response{-1, Msg(CodeServerError)})
		return
	}

	//保存到数据库
	err = service.StoreVideo(user.(User), title, videoUrl, coverUrl)
	if err != nil {
		log.Error("存储视频出现问题：" + err.Error())
		c.JSON(CodeServerError, Response{-1, Msg(CodeServerError)})
		return
	}
	log.Info("上传视频成功")
	c.JSON(CodeSuccess, Response{0, Msg(CodeSuccess)})

}

// 请求： GET，/douyin/publish/list/
// 获取该用户的所有投稿视频
func GetUserVideoList(c *gin.Context) {
	token := c.Query("token")
	user_id := c.Query("user_id")
	if token == "" {
		c.JSON(CodeTokenNotexist, Response{-1, Msg(CodeTokenNotexist)})
		return
	}

	//查数据库，根据userId查出视频列表
	id, _ := strconv.ParseInt(user_id, 10, 64) //转成int64，  参数含义：十进制的64位
	videoList, err := service.GetVideoListByUserId(id)
	if err != nil {
		log.Error(err.Error())
		c.JSON(CodeGetVideoListError, Response{0, Msg(CodeGetVideoListError)})
		return
	}
	//log.Println("视频列表：", videoList)
	c.JSON(CodeSuccess, VideosDto{
		Response:  Response{0, Msg(CodeSuccess)},
		VideoList: videoList,
	})
}
