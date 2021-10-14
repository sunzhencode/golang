FROM golang:1.17.2-alpine3.14 as build
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/http01  .

FROM scratch
COPY --from=build /app/bin/linux/amd64/http01 /app/bin/linux/amd64/http01
ENTRYPOINT ["/app/bin/linux/amd64/http01"]
