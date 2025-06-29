package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(message string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, message)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	msg := os.Getenv("HELLO_MESSAGE")
	if msg == "" {
		msg = "Hello, world!"
	}
	http.HandleFunc("/", helloHandler(msg))
	fmt.Printf("[hello] Listening on :%s ...\n", port)
	http.ListenAndServe(":"+port, nil)
}
