package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)



type Connection struct {
	ws   *websocket.Conn //websocket的连接
	sc   chan []byte     //发送信息的通道
	data  *Data
}

//将http升级为websocket
  //定义升级器, 定义读缓冲区， 写缓冲区， 并设置允许跨域
	var upgrater = &websocket.Upgrader{ReadBufferSize: 1024,
    WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool { return true }}
  //定义myws句柄函数， 进行升级
	func myws(w http.ResponseWriter, r *http.Request){
		//升级websocket, 进行握手
		ws, err := upgrater.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("升级握手失败")
			fmt.Println("err: ", err)
		}
    var c = &Connection{
			ws : ws,
			sc : make(chan []byte, 1024),
			data : &Data{},              
		}


		//子线程向本地写数据
		go c.write()

		c.reader()


	}





//定义读事件
func (c *Connection)reader() {
	for{
		_, m, err := c.ws.ReadMessage()
		//判断客户端是否关闭
		if err != nil {
			fmt.Println("一个客户端关闭了")
			break
		}
		//将网络接收的byte反序列化为结构体
		json.Unmarshal(m, c.data)
		//贩毒案客户端发送的消息
		switch c.data.Type {
		case "login":
			c.Login()
		case "message":
		c.SendMesage()
		case "loginout":
			break
		default:
			fmt.Println("发送的消息有误")
		}

		c.LoginOut()

	

	}

	c.ws.Close()
	//delete(c)
}

//定义写事件
func  (c *Connection)write()  {
	//通过遍历的方式接收管道数据
	for message := range c.sc {
		c.ws.WriteMessage(websocket.TextMessage, message)
	}
	c.ws.Close()
}

//登录
func(c *Connection)Login() {
	h.register <- c
}

//向好友发送信息
func(c *Connection)SendMesage() {
	h.broadcast <- c.data.Content
}

//注销
func(c *Connection)LoginOut() {
	h.u <- c
	c.ws.Close()
}