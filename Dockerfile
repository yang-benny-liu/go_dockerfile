FROM golang AS builder
WORKDIR /go/src/github.com/redhat-developer-demos/go_dockerfile/
COPY ./server.go ./static/. ./go.mod ./
RUN go get -d -v golang.org/x/net/html  
RUN CGO_ENABLED=0 GOOS=linux go build -a -o server .

FROM registry.access.redhat.com/ubi8/ubi-micro
WORKDIR /
COPY --from=builder /go/src/github.com/redhat-developer-demos/go_dockerfile/server /
EXPOSE 8080
CMD ["./server"]
