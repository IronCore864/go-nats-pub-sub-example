FROM golang:latest AS builder
WORKDIR $GOPATH/src/pub
COPY ./go-nats-pub/main.go .
RUN go get ./... && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /pub .
WORKDIR ../sub
COPY ./go-nats-sub/main.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /sub .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /pub .
COPY --from=builder /sub .
ENTRYPOINT ./sub
