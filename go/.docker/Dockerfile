FROM golang:1.14.4 as builder

WORKDIR /go/src/github.com/lgdevlop/k8s-exemplo/go

COPY ./go/hello.go .

COPY ./go/message ./message

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o app hello.go && strip app

FROM scratch

COPY --from=builder /go/src/github.com/lgdevlop/k8s-exemplo/go/app /app

EXPOSE 8000

CMD ["/app"]
