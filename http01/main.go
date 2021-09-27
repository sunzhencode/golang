package main

import (
	"fmt"
	"net/http"
)

func GetMyReuqest(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	for k, v := range h {
		fmt.Println(k, v)
	}
	_, _ = w.Write([]byte("Hello World!"))
}

func main() {
	http.HandleFunc("/", GetMyReuqest)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
	}
}
