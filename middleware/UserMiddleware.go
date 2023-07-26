/*
   @Author Ted
   @Since 2023/7/25 21:35
*/

package middleware

import (
	"douyinProject/common"
	"douyinProject/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 请求用户信息时的中间件：判断token是否存在
func UserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 url里的token参数
		ctx.Query("user_id")
		tokenString := ctx.Query("token")

		fmt.Println("请求token： ", tokenString)

		if tokenString == "" { // !strings.HasPrefix(tokenString, "Bearer "
			ctx.JSON(http.StatusUnauthorized, common.Response{
				-1,
				"权限不足",
			})
			ctx.Abort()
			return
		}

		//tokenString = tokenString[7:] //截取字符，截取”Bearer “之后的内容

		token, claims, err := common.ParseToken(tokenString) //解析token
		fmt.Println("token: ", token)
		fmt.Println("claims: ", claims)

		//如果解析失败，或者解析后token无效，则失败
		if err != nil || !token.Valid {
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
		user, _ := utils.GetUserById(userId)
		fmt.Println("userId: ", userId)
		// 验证用户是否存在
		if user.Id == 0 {
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
