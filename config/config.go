/*
   @Author Ted
   @Since 2023/7/27 11:12
*/

package config

import (
	"github.com/spf13/viper"
	"log"
	"time"
) //Go语言的配置管理库，提供了一种便捷的方式来读取、解析和管理应用程序的配置文件

type Configs struct {
	Mysql MysqlConfig
	Minio MinioConfig
	Path  PathConfig //本地文件base路径的配置
	Level string
	Redis RedisConfig
}

var Config Configs

type PathConfig struct {
	VideoBasePath string
	ImageBasePath string
}
type MysqlConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}
type RedisConfig struct {
	// Network "tcp"
	Network string
	// Addr "127.0.0.1:6379"
	Addr string
	// Password string .If no password then no 'AUTH'. Default ""
	Password string
	// If Database is empty "" then no 'SELECT'. Default ""
	DB string
	// MaxIdle 0 no limit
	MaxIdle int
	// MaxActive 0 no limit
	MaxActive int
	// IdleTimeout  time.Duration(5) * time.Minute
	IdleTimeout time.Duration
	// Prefix "myprefix-for-this-website". Default ""
	Prefix string
}

type MinioConfig struct {
	Host            string
	Port            string
	AccessKeyID     string
	SecretAccessKey string
	Videobuckets    string
	Imagebuckets    string
}

func InitConfig() {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("读取config错误")
	}

	mysql := MysqlConfig{
		Host:     viper.GetString("mysql.host"),
		Port:     viper.GetString("mysql.port"),
		Database: viper.GetString("mysql.database"),
		Username: viper.GetString("mysql.username"),
		Password: viper.GetString("mysql.password"),
	}

	minioConfig := MinioConfig{
		Host:            viper.GetString("minio.Host"),
		Port:            viper.GetString("minio.Port"),
		AccessKeyID:     viper.GetString("minio.AccessKeyID"),
		SecretAccessKey: viper.GetString("minio.SecretAccessKey"),
		Videobuckets:    viper.GetString("minio.Videobuckets"),
		Imagebuckets:    viper.GetString("minio.Imagebuckets"),
	}
	redisConfig := RedisConfig{
		Addr:     viper.GetString("redis.Addr"),
		Password: viper.GetString("redis.Password"),
		DB:       viper.GetString("redis.DB"),
	}
	path := PathConfig{
		VideoBasePath: viper.GetString("minio.VideoBasePath"),
		ImageBasePath: viper.GetString("minio.ImageBasePath"),
	}

	Config = Configs{
		Minio: minioConfig,
		Path:  path,
		Mysql: mysql,
		Redis: redisConfig,
		Level: viper.GetString("level"),
	}
	log.Println("初始化config成功")
}

// 获取config
func GetConfig() Configs {
	return Config
}
