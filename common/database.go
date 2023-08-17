/*
   @Author Ted
   @Since 2023/7/25
*/

package common

import (
	"douyinProject/config"
	"douyinProject/log"
	. "douyinProject/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	mysqlConfig := config.GetConfig().Mysql
	username := mysqlConfig.Username //账号
	password := mysqlConfig.Password //密码
	host := mysqlConfig.Host         //数据库地址，可以是Ip或者域名   ,"127.0.0.1","localhost"
	port := mysqlConfig.Port         //数据库端口
	Dbname := mysqlConfig.Database   //数据库名
	timeout := "10s"                 //连接超时，10秒
	// 临时的 DSN，只用于检查数据库是否存在
	tempDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, timeout)
	tempDB, err := gorm.Open(mysql.Open(tempDSN), &gorm.Config{})
	if err != nil {
		// 处理连接错误
		panic(err)
	}
	tempDB.Exec("CREATE DATABASE IF NOT EXISTS douyin")
	tmp, _ := tempDB.DB()
	tmp.Close()
	if err != nil {
		// 处理关闭错误
		panic(err)
	}

	//数据库连接的 DSN（Data Source Name），其中包含了数据库连接的相关信息
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			SkipDefaultTransaction: true,  //关闭默认事务，性能优化
			PrepareStmt:            true}) //缓存预编译语句，提高35%左右性能
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 连接成功
	log.Info("db连接成功")
	// 判断数据库是否存在,不存在则创建
	return db
}

var DB *gorm.DB

func DBInit() {
	DB = Connection()
	DB.AutoMigrate(&User{}, &Video{}, &Favorite{}, &Relation{}, &Comment{}) //,
	log.Info("数据库初始化成功")
}

func GetDB() *gorm.DB {
	return DB
}
