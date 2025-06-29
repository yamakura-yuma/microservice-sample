package handler

import (
	"fmt"
	"net/http"
)

func HelloHandler(message string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, message)
	}
}
