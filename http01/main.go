package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// “/”路径函数
func GetMyRequest(w http.ResponseWriter, r *http.Request) {
	//获取全部请求头
	h := r.Header
	//从请求携带的参数中获取想要获取的环境变量，并返回到reqponse header
	query := r.URL.Query()
	env := query.Get("env")
	version := os.Getenv(env)
	fmt.Println(version)
	if version != "" {
		w.Header().Set("version", version)
	}
	//将获取的请求头循环返回到reqponse header
	for k, v := range h {
		//fmt.Println(k, v)
		value := strings.Join(v, ",")
		w.Header().Set(k, value)
	}
	_, _ = w.Write([]byte("Hello World!"))
}

//健康检查函数
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/", GetMyRequest)
	http.HandleFunc("/healthz", Healthz)
	if err := http.ListenAndServe("0.0.0.0:80", nil); err != nil {
		fmt.Println(err)
	}
	//err := http.ListenAndServe("0.0.0.0:80", nil)
	//if err != nil {
	//	fmt.Println(err)
	//}
}
