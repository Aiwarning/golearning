// main.go
package main

/*
流程图:

1. 用户认证系统

   1.1 定义 User 模型
       - 结构体：User
         - 字段：ID, Username, Password

   1.2 初始化数据库
       - 使用 GORM 连接 SQLite 数据库

   1.3 创建 Gin 路由
       - 初始化 Gin 引擎
       - 创建注册和登录路由

   1.4 注册处理程序
       - 路由：/register
       - 方法：register(c *gin.Context)
         - 从请求中获取 JSON 数据
         - 创建新用户并存储在数据库中
         - 返回注册成功的消息

   1.5 登录处理程序
       - 路由：/login
       - 方法：login(c *gin.Context)
         - 从请求中获取 JSON 数据
         - 通过用户名查找用户
         - 检查密码
         - 返回登录成功的消息
*/

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

// User 模型
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex"`
	Password string `json:"password"`
}

var db *gorm.DB
var err error

func main() {
	// 初始化数据库
	/*
		gorm.Open: 这是 GORM 提供的用于打开数据库连接的函数。它接受两个参数：数据库驱动和数据库连接配置
		sqlite.Open("test.db"): 这部分指定了数据库驱动和连接信息。在这里，我们使用了 SQLite 驱动，然后指定数据库文件名为 "test.db"。
								如果你使用其他数据库，例如 MySQL、PostgreSQL，你需要相应地更改驱动和连接信息。
		&gorm.Config{}: 这是 GORM 的配置选项，它是一个结构体，用于配置数据库连接的一些选项。在这里，我们使用了默认配置。

	*/
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	/*	db.AutoMigrate 是 GORM（Go的对象关系映射库）提供的一个方法，用于自动创建或更新数据库表结构，以便与模型定义保持一致。
		&User{}： 这是一个 GORM 模型，通常是一个结构体，表示数据库中的一张表。在这个例子中，我们使用 User 结构体，它可能表示用户表
		执行数据库迁移：
		db.AutoMigrate(&User{}) 表示执行数据库迁移，确保数据库中有与 User 模型对应的表结构。
		如果表不存在，它将创建该表；如果表已经存在，它将检查表结构是否与模型定义一致，如果不一致则进行更新。
	*/
	db.AutoMigrate(&User{})

	// 初始化 Gin 路由
	r := gin.Default()

	// 路由
	r.POST("/register", register)
	r.POST("/login", login)

	// 运行服务器
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

// 注册处理程序
func register(c *gin.Context) {
	/*
		用途不同：
				这种方式更符合 Go 语言的设计哲学，鼓励显式的类型声明
				type 用于创建新的类型，起到类型定义的作用。
				var 用于声明变量，起到变量定义的作用
			// 使用 type 创建新类型
			type Celsius float64

			// 使用 var 声明变量
			var temperature Celsius

	*/
	var user User

	/*
		使用 ShouldBindJSON 方法从请求中解析 JSON 数据，并将解析后的数据绑定到 user 变量。
	*/
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 创建新用户
	db.Create(&user)

	c.JSON(200, gin.H{"message": "用户注册成功"})
}

// 登录处理程序
func login(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 通过用户名查找用户
	result := db.Where("username = ?", user.Username).First(&user)
	if result.Error != nil {
		c.JSON(401, gin.H{"error": "凭据无效"})
		return
	}

	// 检查密码
	if user.Password != user.Password {
		c.JSON(401, gin.H{"error": "凭据无效"})
		return
	}

	c.JSON(200, gin.H{"message": "登录成功"})
}
