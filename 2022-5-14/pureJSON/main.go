package main

import "github.com/gin-gonic/gin"

/***
调用这个两个api后，我们可以发现他们不一样的地方
1. /json
{"html":"\u003cb\u003eHello, world!\u003c/b\u003e"}
2./porejson
{"html":"Hello, world!"}
我们查看了文档后，我们发现了，第一个是没有将标签进行解析的，也就是说这里返回的是： unicode 替换特殊 HTML 字符
然后第二个就是直接对html的标签进行了解析的。
*/
func main() {
	r := gin.Default()

	// 提供 unicode 实体
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 提供字面字符
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
