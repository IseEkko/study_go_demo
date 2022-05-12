package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

/***
学习内容：websocket
  什么是websocket,对于这个东西，我的理解就是增强版的http，是长链接的，是一个tcp链接
在之前要达到一个实时的消息发送，我们需要做的是轮询的方式，
*/
//创建websockt对象
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// remove the previous fmt statement
	// fmt.Fprintf(w, "Hello World")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	//返回一个链接对象
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")
	//返回的数据，返回给前端
	//参数是一个类型，一个[]byte数组
	err = ws.WriteMessage(1, []byte("Hi Client ceshi!"))
	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	go reader(ws)
}

//读数据
func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		p = append(p, 'f')
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	//定义websockt的读写缓冲

	fmt.Println("Hello World")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
