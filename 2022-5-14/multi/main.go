package main

import (
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

/**
获取表单的数据，这里绑定的是结构体
这里展示来两个方法，第一个c.ShouldBindWith(&form, binding.Form)，使用这方法，那么我们最后只能获取到form-data的数据
第二个方式： c.ShouldBind(&form)
这个方式可以自己选定获取的方式。
*/
func main() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		// 你可以使用显式绑定声明绑定 multipart form：
		// c.ShouldBindWith(&form, binding.Form)
		// 或者简单地使用 ShouldBind 方法自动绑定：
		var form LoginForm
		// 在这种情况下，将自动选择合适的绑定
		if c.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		} else {
			c.JSON(422, gin.H{
				"err": c.ShouldBind(c.ShouldBind(&form)),
				"msg": "参数错误",
			})
		}
	})
	router.Run(":8080")
}
