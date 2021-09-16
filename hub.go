package main

import (
	"encoding/json"
	"fmt"
)

type hub struct {
	c          map[*Connection]bool
	register   chan *Connection
	broadcast  chan []byte
	unregister chan *Connection
}

var h = hub{
	c:          make(map[*Connection]bool),
	register:   make(chan *Connection),
	unregister: make(chan *Connection),
	broadcast:  make(chan []byte),
}

func (h *hub) run() {

	for {
		select {
		case conn := <-h.register:
			h.RegUser(conn)
		case conn := <-h.unregister:
			h.UnReg(conn)
		case data_b := <-h.broadcast:
			h.BroadCast(data_b)
		default:
			fmt.Println("hub:是不是发错了啊！")
		}
	}

}

//注册新用户
func (h *hub) RegUser(conn *Connection) {
	h.c[conn] = true
	conn.data.Type = "handshake"
	data, _ := json.Marshal(conn.data)
	h.broadcast <- data
}

//注销用户
func (h *hub) UnReg(conn *Connection) {
	delete(h.c, conn)

}

//广播消息
func (h *hub) BroadCast(data []byte) {
	for conn := range h.c {
		select {
		case conn.sc <- data:
		default:
			conn.LoginOut()
			h.UnReg(conn)
		}
	}

}
