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
```