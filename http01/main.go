package main

import (
	"fmt"
	"http01/serverfunc"
	"net/http"
)

func main() {
	http.HandleFunc("/", serverfunc.GetMyRequest)
	http.HandleFunc("/healthz", serverfunc.Healthz)
	if err := http.ListenAndServe("0.0.0.0:80", nil); err != nil {
		fmt.Println(err)
	}
	//err := http.ListenAndServe("0.0.0.0:80", nil)
	//if err != nil {
	//	fmt.Println(err)
	//}
}
