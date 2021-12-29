## use istio ingress gateway
### install istio
```bash
curl -L https://istio.io/downloadIstio | sh -
cd istio-1.12.0
cp bin/istioctl /usr/local/bin
istioctl install --set profile=demo -y
```

### create secret
```bash
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=cncamp Inc./CN=www.example.com' -keyout example.key -out example.crt
kubectl create -n istio-system secret tls example-credential --key=example.key --cert=example.crt
```

### deploy svc
```bash
kubectl create ns istiosvc
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
kubectl apply -f istio-spec.yaml -n istiosvc
```

### test
```bash
curl --resolve www.example.com:443:$INGRESS_IP https://www.example.com/healthz -v -k


[root@sunzhen istio]# curl --resolve www.example.com:443:$INGRESS_IP https://www.example.com/  -k
Hello World!  
```

### L7 route
```bash
kubectl apply -f istio-l7.yaml -n istiosvc

[root@sunzhen istio]# curl --resolve www.example.com:443:$INGRESS_IP https://www.example.com/httpserver  -k
Hello World!
```

## tracing
### edit main.go
```
req, err := http.NewRequest("GET", "http://service1", nil)
if err != nil {
fmt.Printf("%s", err)
}
// 在go语言中需要把haader转化为小写，才可以向下正确传递
lowerCaseHeader := make(http.Header)
for key, value := range r.Header {
lowerCaseHeader[strings.ToLower(key)] = value
}
glog.V(4).Info("headers:", lowerCaseHeader)
req.Header = lowerCaseHeader
client := &http.Client{}
resp, err := client.Do(req)

```

### install jaeger
```bash
kubectl apply -f jaeger.yaml
kubectl edit configmap istio -n istio-system

      tracing:
        zipkin:
          address: zipkin.istio-system:9411
        sampling: 100
```
deploy tracing
```bash
kubectl create ns tracing
kubectl label ns tracing istio-injection=enabled
kubectl -n tracing apply -f service0.yaml
kubectl -n tracing apply -f service1.yaml
kubectl -n tracing apply -f service2.yaml 
kubectl apply -f istio-specs.yaml -n tracing
```


### forward jaeger
```bash
istioctl dashboard jaeger --address 0.0.0.0

click on Find Traces
```