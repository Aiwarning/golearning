package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func gintest() {
	// 创建 Gin 引擎
	router := gin.Default()

	// 设置 session 中间件，使用 cookie 存储
	// store := cookie.NewStore([]byte("secret"))用于创建一个基于 cookie 的会话存储。
	//	cookie.Store 实现了 sessions.Store 接口，它负责存储和检索会话数据
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	/*
		路由注册： router.GET("/set", ...) 注册了一个处理 GET 请求的路由，该路由的路径是 "/set"。
		处理函数： func(c *gin.Context) 是处理请求的函数。在这个函数中，我们使用了 gin.Context 类型的参数 c，它包含了关于当前 HTTP 请求的信息和方法
		Session 设置：	session := sessions.Default(c) 创建	了一个与当前请求关联的 Session 对象。
									sessions.Default 用于获取默认的 Session 对象，而具体的 Session 存储（例如 cookie 存储）会与 Gin 的中间件一起使用
		session.Set("username", "example_user")：设置了 Session 中键名为 "username" 的值为 "example_user"。这里用来模拟在用户登录成功后，将用户名存储在 Session 中。
		session.Save() 将修改后的 Session 保存，确保值被持久化
		c.JSON(200, gin.H{"message": "Session set successfully"})
	*/
	router.GET("/set", func(c *gin.Context) {
		// 在 session 中设置值
		session := sessions.Default(c)
		session.Set("username", "example_user")
		session.Save()

		c.JSON(200, gin.H{"message": "Session set successfully"})
	})

	router.GET("/get", func(c *gin.Context) {
		// 从 session 中获取值
		session := sessions.Default(c)
		username := session.Get("username")

		if username == nil {
			c.JSON(200, gin.H{"message": "No username in session"})
		} else {
			c.JSON(200, gin.H{"username": username})
		}
	})

	// 启动服务器
	router.Run(":8080")
}
