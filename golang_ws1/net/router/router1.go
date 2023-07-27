package router

import (
	"fmt"
	"net/http"
)

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		name = "go lang http"
	}
	msg := []byte(fmt.Sprintf("hello %v", name))
	writer.Write(msg)
}

func RegistHandlers() {
	http.HandleFunc("/hello", helloHandler)
}
