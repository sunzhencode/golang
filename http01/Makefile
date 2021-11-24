export tag=v3.0
build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

docker-build:
	docker build -t sunzhenreg/golang-http:${tag} .
	docker push sunzhenreg/golang-http:${tag}