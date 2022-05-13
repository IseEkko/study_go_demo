package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/***
演示go渲染喊他们html模版demo，这里演示了如何使用gin进行渲染，这里需要注意的就是关于路径的书写
LoadHTMLGlob函数的路径，是从启动的路径开始计算的。
HTML函数，第一个参数传入状态，然后传入渲染文件名称，也就是说上面的函数，就是把文件加载进来，然后在这里调用，
这里有一个细节，就是调用的名称是可以自己定义的，使用的格式：
{{ define "posts/index.tmpl" }}
<html><h1>
	{{ .title }}
</h1>
<p>Using posts/index.tmpl</p>
</html>
{{ end }}
第三个参数就是进行的参数传送，使用的是json。
*/
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("2022-5-13/view/*")
	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"titile": "ceshi",
		})
	})
	router.Run(":8080")
}
