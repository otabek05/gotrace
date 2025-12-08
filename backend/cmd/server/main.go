package main

import (
	"gotrace/backend/internal/api"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)


func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}


func main() {
	r := gin.Default()
	api.RegisterRoutes(r)
	log.Println("Server running in port: 8081")
	if err := http.ListenAndServe(":8081", enableCORS(r)); err != nil {
		log.Fatal(err)
	}
}