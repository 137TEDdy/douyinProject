/*
   @Author Ted
   @Since 2023/7/25
*/

package common

import (
	. "douyinProject/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	username := "root"          //账号
	password := "gjh1375809659" //密码
	host := "127.0.0.1"         //数据库地址，可以是Ip或者域名   ,"127.0.0.1","localhost"
	port := 3306                //数据库端口
	Dbname := "douyin"          //数据库名
	timeout := "10s"            //连接超时，10秒
	//数据库连接的 DSN（Data Source Name），其中包含了数据库连接的相关信息
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 连接成功
	fmt.Println("连接成功， ", db)
	return db
}

func DBInit() {
	db := GetDB()
	db.AutoMigrate(&User{}, &Video{}, &Favorite{}, &Relation{}, &Comment{}) //,
	fmt.Println("数据库创建成功")
}

func GetDB() *gorm.DB {
	var db = Connection()
	return db
}

var DB *gorm.DB = GetDB()
