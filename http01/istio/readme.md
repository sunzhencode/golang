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
lowerCaseHeader := make(http.Header)
for key, value := range r.Header {
lowerCaseHeader[strings.ToLower(key)] = value
}
glog.V(4).Info("headers:", lowerCaseHeader)
req.Header = lowerCaseHeader
client := &http.Client{}
resp, err := client.Do(req)

```
