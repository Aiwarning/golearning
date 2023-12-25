// main.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 模拟用户数据
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func main() {
	router := gin.Default()

	// 使用登录校验中间件
	router.Use(AuthMiddleware())

	// 定义一个需要登录的路由
	router.GET("/protected", func(c *gin.Context) {
		user := c.MustGet("user").(string)
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Hello, %s!", user)})
	})

	// 启动服务
	router.Run(":8080")
}

// AuthMiddleware 登录校验中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户身份信息，这里假设通过请求头或者其他方式传递用户信息
		user := c.GetHeader("Authorization")

		// 模拟简单的登录校验
		if password, ok := users[user]; ok {
			// 将用户信息传递给后续处理函数
			c.Set("user", user)
			c.Next()
		} else {
			// 用户未登录，返回未经授权的错误
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
		}
	}
}
