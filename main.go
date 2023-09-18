package main

import (
	"fmt"
	"go-learn/router"
	"log"
	"net/http"
	"os"
)

func main() {
	handler := router.New()
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")),
		Handler: handler,
	}
	log.Println("server running at", server.Addr)

	server.ListenAndServe()

}
