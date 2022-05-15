package main

import "github.com/gin-gonic/gin"

/***
获取参数
*/
func main() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		//表单参数
		message := c.PostForm("message")
		//这里是设置字段的默认值
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8080")
}
