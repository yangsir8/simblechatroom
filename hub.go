package main

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
