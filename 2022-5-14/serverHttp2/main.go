package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

/***
使用服务推送功能，是http2的功能，在go1.8的时候是可以使用的
在传统的http请求中，一个请求只能带一个资源，那么这个时候就会出现一个问题，当一个页面有很多资源的时候，我们需要一个一个的区请求
服务端接收到一个请求给一个资源
我们知道的一个http请求的本质是tcp，那么这么多的tcp请求肯定是对资源的浪费，那么我们如何来解决这个问题呢？
第一个：
我们可以使用preload技术，将资源提前加载进这个地方。这样我们也可以对性能进行提升，但是还是不是最优的，
第二个：
也就是我们今天学习的功能，http2的服务推送，他在请求这个的时候，我们将资源一起推送给他。那么就是一个请求可以得到多个回复
*/
func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)

	r.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// 使用 pusher.Push() 做服务器推送
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})

	// 监听并在 https://127.0.0.1:8080 上启动服务
	r.Run(":8080")
}
