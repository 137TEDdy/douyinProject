/*
   @Author Ted
   @Since 2023/7/25 21:35
*/

package middleware

import (
	"douyinProject/common"
	"douyinProject/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 请求用户信息时的中间件：判断token是否存在
func TokenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//调试
		//log.Println(ctx.ContentType())
		//log.Println("form参数: ", ctx.Request.PostForm)

		//处理两种传入token的情况,更具广泛性
		var tokenString string
		tokenString = ctx.PostForm("token")
		if tokenString == "" {
			//log.Println("从query里获取token")
			tokenString = ctx.Query("token")
		}
		//log.Println("请求token： ", tokenString)

		if tokenString == "" { // !strings.HasPrefix(tokenString, "Bearer "
			log.Println("权限不足")
			ctx.JSON(http.StatusUnauthorized, common.Response{
				-1,
				"权限不足",
			})
			ctx.Abort()
			return
		}

		//tokenString = tokenString[7:] //截取字符，截取”Bearer “之后的内容
		token, claims, err := common.ParseToken(tokenString) //解析token

		//如果解析失败，或者解析后token无效，则失败
		if err != nil || !token.Valid {
			log.Println("权限不足")
			ctx.JSON(http.StatusUnauthorized, common.Response{
				-1,
				"权限不足",
			})
			ctx.Abort()
			return
		}

		//此时token通过验证, 我们可以获取claims中的UserID
		userId := claims.UserId
		//根据id获取user
		user, _ := service.GetUserById(userId)
		// 验证用户是否存在
		if user.Id == 0 {
			log.Println("权限不足")
			ctx.JSON(http.StatusUnauthorized, common.Response{
				-1,
				"权限不足",
			})
			ctx.Abort()
			return
		}

		//用户存在 将user信息写入上下文,并放行
		ctx.Set("user", user)
		ctx.Next()
	}
}
