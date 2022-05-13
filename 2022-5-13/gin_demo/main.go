package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/***
这里演示了一下，gin的渲染模式，这里我们需要注意的事情，我们启动的时候，其实是在study_demo下的。

我们可以给文件使用{define "users/index.tmpl"}给文件命名

需要注意的事情，我们在加载了templates后要想渲染post下的文件，我们需要给他命名define "users/index.tmpl"}
这样我们后面才能加载出来。
*/
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("2022-5-13/gin_demo/templates/**/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")
}
