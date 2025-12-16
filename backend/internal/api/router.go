package api

import (
	"gotracer/internal/ws"
	"net/http"
)


func NewRouter() http.Handler {
	mux := http.NewServeMux()
	go ws.DefaultHub.Run()
		
	mux.HandleFunc("/ws", ws.DefaultHub.ServeWS)
	mux.HandleFunc("/api/v1/network/interfaces", NetInterfaceHandler)

	mux.Handle("/", spaHandler("../../build/web"))

	return withCORS(mux)
}

