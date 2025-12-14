package api

import (
	"gotracer/internal/ws"
	"net/http"
)


func NewRouter() http.Handler {
	mux := http.NewServeMux()
	go ws.DefaultHub.Run()
		
	mux.HandleFunc("/ws", ws.DefaultHub.ServeWS)
	mux.HandleFunc("/api/v1/interfaces", NetInterfaceHandler)
	return mux 
}