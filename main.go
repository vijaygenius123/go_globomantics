package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct {
	Message string
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bar Called"))
}

func main() {
	fmt.Println("Hello World")
	http.Handle("/foo", &fooHandler{Message: "Hello"})
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":5000", nil)
}
