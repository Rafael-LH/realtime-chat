package main

import (
	"fmt"
	"net/http"

	"github.com/EddCode/realtime-chat/middelware"
)

func setUpRoutes() {
	http.HandleFunc("/", middelware.Home)
	http.HandleFunc("/ws", middelware.ServerWs)
}

func main() {
	fmt.Println("Chat App v0.01")
	setUpRoutes()
	http.ListenAndServe(":5000", nil)
}
