### 作业 编写一个 HTTP 服务器

1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 当访问 localhost/healthz 时，应返回 200

### 使用方法

* linux下运行文件 bin/linux/amd64/http01
* windows下运行文件 bin/windows/amd64/http01.exe 

### 测试方法
* 访问根路径并且获取环境变量，可以在响应头中查看到内容，并且可以在控制台看到访问日志
```bash
curl -I 127.0.0.1?env=VERSION
curl -I 127.0.0.1?env=GOROOT 
```
* 健康检查
```bash
curl -I 127.0.0.1/health
```

