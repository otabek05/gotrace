package main

import (
	"fmt"
	"gotracer/internal/api"
	"log"
	"net/http"
)


func main() {

	router := api.NewRouter()
	fmt.Println("Server is starting on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}