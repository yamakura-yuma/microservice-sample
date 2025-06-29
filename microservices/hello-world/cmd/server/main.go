package main

import (
	"fmt"
	"microservices/hello-world-server/internal/handler"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	msg := os.Getenv("HELLO_MESSAGE")
	if msg == "" {
		msg = "Hello, world!"
	}
	http.HandleFunc("/", handler.HelloHandler(msg))
	fmt.Printf("[hello] Listening on :%s ...\n", port)
	http.ListenAndServe(":"+port, nil)
}
