package serverfunc

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
)

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
		w.Header().Set(k, value)
	}
	w.WriteHeader(200)
	_, _ = w.Write([]byte("Hello World!"))
	LogPrint(r)
}

// Healthz 作业四：健康检查函数
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	LogPrint(r)
}

// LogPrint 作业三：使用生产者消费者模型，打印日志的函数
func LogPrint(r *http.Request) {
	// 获取客户端IP
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	AccessLog := fmt.Sprintf("clientip:%s,ststuscode:%d,URL:%s", ip, 200, r.URL)
	q := Queue{
		queue: []string{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}
	go q.SendLog(AccessLog)    // 向生产者发送数据
	go fmt.Println(q.GetLog()) //从消费者中获取数据
}
