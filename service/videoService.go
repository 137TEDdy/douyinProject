/*
   @Author Ted
   @Since 2023/7/25 19:30
*/

package service

import (
	"douyinProject/config"
	"douyinProject/repo"
	"douyinProject/utils"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)
import . "douyinProject/model"

func GetVideoList(user_id int64) ([]*Video, error) {
	var videoList []*Video
	var err error
	if user_id != 0 { //如果用户已经登录
		videoList, err = repo.GetVideoListLogin(user_id)
	} else { //用户未登录时
		videoList, err = repo.GetVideoListUnLogin()
	}

	return videoList, err
}

// 根据用户id查出用户，以及对应的视频
func GetVideoListByUserId(userId int64) ([]*Video, error) {
	videolist, err := repo.GetVideoListByUserID(userId)
	return videolist, err
}

func StoreVideo(user User, title, videoUrl, coverUrl string) error {
	video := Video{
		Author:      user,
		Title:       title,
		AuthorId:    user.Id,
		PlayUrl:     videoUrl,
		CoverUrl:    coverUrl,
		PublishTime: utils.GetCurrentTime(),
	}
	err := repo.StoreVideo(video)
	return err
}

// 使用ffmpeg，截取视频的第一秒作为视频封面，并存储到outputpath里
func GetImage(videoPath string) string {
	imageBasePath := config.GetConfig().Path.ImageBasePath
	arrs := strings.Split(videoPath, "\\") //分割视频路径，取出视频名
	lastName := arrs[len(arrs)-1]
	lastName = lastName[:len(lastName)-4] + ".jpg" //获取  xxx.mp4里的xxx，变成xxx.jpg

	outputPath := filepath.Join(imageBasePath, lastName) // D:\2\xxx.jpg
	log.Println("VideoPath: ", videoPath, "   ouputPath: ", outputPath)
	// 指定FFmpeg命令和参数，截取1s处图片
	cmd := exec.Command("D:\\APP2\\ffmpeg-5.1.2-essentials_build\\bin\\ffmpeg", "-i", videoPath, "-ss", "1", "-f", "image2", "-t", "0.01", "-y", outputPath)
	//cmd := exec.Command("D:\\APP2\\ffmpeg-5.1.2-essentials_build\\bin\\ffmpeg", "-i", videoPath, "-ss", "00:00:01", "-vframes", "1", "-y", outputPath) 这个无法执行

	// 执行FFmpeg命令
	err := cmd.Run()
	if err != nil {
		log.Fatal("ffmpeg出现错误:", err)
	}
	log.Println("视频截图已保存到", outputPath)
	return outputPath
}
