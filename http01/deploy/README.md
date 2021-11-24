* 部署deployment
```text
kubectl apply -f deployment.yaml


replicas: 2  副本数设置为2，保证服务高可用
健康检查接口：/healthz 
使用pod反亲和，使pod部署在不同的主机上
设置requests 小于 limit，pod为Burstable等级，同时提高了node资源利用率
程序读取env环境变量配置监听地址和端口，配置与代码分离  SERVER_ADDR:SERVER_PORT 默认：0.0.0.0:80
通过启动参数flag设置glog的日志级别和日志保存目录，日志保存在vulome 的emptyDir中
```

* 部署service
```text
kubectl apply -f service.yaml
```

* 部署ingress 提供访问入口
```text
使用准备好的证书生成tls secret
kubectl create secret tls example-tls --key keyfile --cert crtfile -n default

kubectl apply -f ingress.yaml
```