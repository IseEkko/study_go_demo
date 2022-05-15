package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//参数获取，get post，包含定义默认值的操作
//关于psot的默认值在前面是提到了的
func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	router.Run(":8080")
}
