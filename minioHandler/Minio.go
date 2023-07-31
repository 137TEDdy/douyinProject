/*
   @Author Ted
   @Since 2023/7/27 10:18
*/

package minioHandler

import (
	"douyinProject/config"
	"douyinProject/utils"
	"github.com/minio/minio-go/v6"
	"github.com/sirupsen/logrus"
	"log"
	"strconv"
	"strings"
)

// 封装minio的各种属性
type Minio struct {
	MinioClient   *minio.Client
	endpoint      string
	port          string
	VideoBuckets  string
	Imagebuckets  string
	VideoBasePath string
}

var Client Minio

func GetClient() Minio {
	return Client
}

// 初始化minio
func InitMinio() {
	conf := config.GetConfig() //读取配置
	endpoint := conf.Minio.Host
	port := conf.Minio.Port
	endpoint = endpoint + ":" + port      //例如：192.168.88.131:9000
	accessKeyID := conf.Minio.AccessKeyID //登录时使用
	secretAccessKey := conf.Minio.SecretAccessKey
	videoBucket := conf.Minio.Videobuckets
	imagesBucket := conf.Minio.Imagebuckets
	videoBasePath := conf.Path.VideoBasePath //视频本地存储base路径
	useSSL := false

	// 初使化 minio client对象， 创建一个客户端
	log.Println("链接地址：", endpoint)
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		logrus.Error("创建minio客户端错误: ", err)
	}
	//创建存储桶
	creatBucket(minioClient, imagesBucket)
	creatBucket(minioClient, videoBucket)
	Client = Minio{minioClient, endpoint, port, videoBucket, imagesBucket, videoBasePath}
}

// 创建minio的桶（类似数据库的表）
func creatBucket(m *minio.Client, bucketName string) {

	isExist, err := m.BucketExists(bucketName) //判断bucket是否存在
	if err != nil {
		logrus.Errorf("检查桶时出错，check %s bucketExists err:%s", bucketName, err.Error())
	}
	if !isExist {
		m.MakeBucket(bucketName, "us-east-1") //不存在就创建，us-east-1为地区、
		log.Println(bucketName, "桶不存在，初始化")
	}
	log.Println(bucketName, "桶已存在")
	//设置桶策略，允许任何AWS账号执行 s3:GetObject 操作， 授予公共读取权限
	policy := `{"Version": "2012-10-17",
				"Statement": 
					[{
						"Action":["s3:GetObject"],     
						"Effect": "Allow",
						"Principal": {"AWS": ["*"]},
						"Resource": ["arn:aws:s3:::` + bucketName + `/*"],
						"Sid": ""
					}]
				}`
	err = m.SetBucketPolicy(bucketName, policy) //设置桶权限
	if err != nil {
		logrus.Errorf("SetBucketPolicy %s  err:%s", bucketName, err.Error())
	}
}

// 上传文件
func (m *Minio) UploadFile(userID int64, filetype, filePath string) (string, error) {
	log.Println("Minio的上传文件函数")
	var fileName strings.Builder //string不可变，引入builder能够高效地构建和修改字符串
	var contentType, suffix, bucketName string

	//判断传入的是视频还是图片，设置对应参数
	if filetype == "video" {
		contentType = "video/mp4"
		suffix = ".mp4"
		bucketName = m.VideoBuckets
	} else if filetype == "image" {
		contentType = "image/jpeg"
		suffix = ".jpg"
		bucketName = m.Imagebuckets
	}

	fileName.WriteString(strconv.Itoa(int(userID)))
	fileName.WriteString("_")
	fileName.WriteString(strconv.FormatInt(utils.GetCurrentTime(), 10)) //int转换成字符串，十进制表示
	fileName.WriteString(suffix)
	log.Println("filename： ", fileName.String())
	log.Println("bucketname： ", bucketName)
	log.Println("文件路径filepath： ", filePath)

	n, err := m.MinioClient.FPutObject(bucketName, fileName.String(), filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		log.Println("更新文件错误:%s", err.Error())
		return "", err
	}
	log.Println("更新 %dbyte大小的文件成功，文件名:%s", n, fileName)

	url := "http://" + m.endpoint + "/" + bucketName + "/" + fileName.String()
	return url, nil
}
