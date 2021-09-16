package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()
	fmt.Println("ok, 正在监听")
	router.HandleFunc("/ws", myws)
	if err := http.ListenAndServe("127.0.0.1:8000", router); err != nil{
		fmt.Println("服务器监听失败")
		return
	}
	fmt.Println("ok, 正在监听")

}