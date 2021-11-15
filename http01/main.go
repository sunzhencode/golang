package main

import (
	"fmt"
	"github.com/golang/glog"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	//flag.Set("v", "4")
	//flag.Parse()
	log.Println("start server")
	http.HandleFunc("/", GetMyRequest)
	http.HandleFunc("/healthz", Healthz)
	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// GetMyRequest “/”路径函数
func GetMyRequest(w http.ResponseWriter, r *http.Request) {
	//获取全部请求头
	h := r.Header
	//作业二：从请求携带的参数中获取想要获取的环境变量，并返回到reqponse header
	query := r.URL.Query()
	env := query.Get("env")
	version := os.Getenv(env)
	if version != "" {
		w.Header().Set("version", version)
	}
	//作业一：将获取的请求头循环返回到reqponse header
	for k, v := range h {
		//fmt.Println(k, v)
		value := strings.Join(v, ",")
		w.Header().Set(k, value) //写入到ResponseHeader
		//io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v)) //返回到页面
	}
	w.WriteHeader(200)
	_, _ = w.Write([]byte("Hello World!"))
	glog.V(4).Info(logFmt(r))
}

// Healthz 健康检查函数
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	glog.V(4).Info(logFmt(r))
}

// 日志格式化 并获取真实IP
func logFmt(r *http.Request) string {
	// 获取客户端IP
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	AccessLog := fmt.Sprintf("clientip:%s,ststuscode:%d,URL:%s", ip, 200, r.URL)
	return AccessLog
}
