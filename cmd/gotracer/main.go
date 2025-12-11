package main

import (
	"fmt"
	"gotrace/internal/ws"
	"net/http"
)


func main() {
	go ws.DefaultHub.Run()

	http.HandleFunc("/ws", ws.DefaultHub.ServeWS)

	fmt.Println("Server is running on port : 8080")
	http.ListenAndServe(":8080", nil)
}