### 作业 编写一个 HTTP 服务器

1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 当访问 localhost/healthz 时，应返回 200

### 部署要求
1. 优雅启动
2. 优雅终止
3. 资源需求和 QoS 保证
4. 探活
5. 日常运维需求，日志等级
6. 配置和代码分离

### 目录结构
* 代码 main.go
* kubernetes部署yaml: deploy目录下

### 编译二进制文件
```bash
make build
```
### 构建镜像
```bash
make docker-build
```

### kubernetes 部署说明
```txt
使用健康检查接口：/healthz 探测应用状态，实现pod声明周期管理
使用pod反亲和，使pod部署在不同的主机上
设置requests 小于 limit，pod为Burstable等级，同时提高了node资源利用率
使用env配置启动地址和端口，配置与代码分离  SERVER_ADDR:SERVER_PORT 默认：0.0.0.0:80
```

### 命令行启动服务
```bash
go run main.go -log_dir log -alsologtostderr -v 4
-log_dir 日志输出的路径
-alsologtostderr 日志同时输出到控制台
-v 4 日志级别
```

### 测试方法
* 访问根路径并且获取环境变量，可以在响应头中查看到内容，并且可以在控制台看到访问日志
```bash
curl -I 127.0.0.1?env=version 
curl -I 127.0.0.1?env=GOROOT 
```
* 健康检查
```bash
curl -I 127.0.0.1/health
```