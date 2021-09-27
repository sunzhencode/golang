package main

import (
	"fmt"
	"net/http"
	"strings"
)

func GetMyReuqest(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	for k, v := range h {
		fmt.Println(k, v)
		value := strings.Join(v, ",")
		w.Header().Set(k, value)
	}
	_, _ = w.Write([]byte("Hello World!"))
}

func main() {
	http.HandleFunc("/", GetMyReuqest)
	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		fmt.Println(err)
	}
}
