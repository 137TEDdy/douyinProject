/*
   @Author Ted
   @Since 2023/7/25 21:04
*/

package common

import (
	"douyinProject/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 定义jwt加密密钥
var jwtKey = []byte("a_secret_crect") //不能单引号，必须双引号

//自定义的Claims结构体，用于存储JWT的载荷信息。
//Claims结构体包含一个UserId字段和jwt.StandardClaims字段，后者是JWT包中提供的标准声明。

type Claims struct {
	UserId int64
	jwt.StandardClaims
}

// 创建一个带有用户特定信息的JWT，并对其进行签名，生成一个JWT字符串用于身份验证和授权。
func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) //过期时间：七天
	claims := &Claims{
		UserId: user.Id,
		//ExpiresAt（过期时间）、IssuedAt（签发时间）、Issuer（签发者）和Subject（主题）
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "jkdev.cn",
			Subject:   "user token",
		},
	}
	//创建了一个新的JWT对象token，使用HS256算法进行签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//使用密钥jwtKey对JWT进行签名，并将签名后的JWT字符串赋给tokenString
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 用于解析JWT字符串
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	//解析JWT字符串。它接受三个参数：tokenString表示待解析的JWT字符串，claims表示用于存储解析出的声明信息的空结构体实例，
	//最后一个参数是一个回调函数，用于提供用于验证签名的密钥。在这个回调函数中，它简单地返回之前定义的jwtKey作为密钥。
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
