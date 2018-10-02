FROM golang:1.11 as builder
WORKDIR /go/src/github.com/emeve89/microgo/consignment-service
COPY . .
RUN go get github.com/golang/dep/cmd/dep && dep init && dep ensure && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM debian:latest
RUN apk --no-cache add ca-certificates && mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/emeve89/microgo/consignment-service/consignment-service .
CMD ["./consignment-service"]