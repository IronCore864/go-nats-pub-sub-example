# NATS Simple Pub/Sub Usage with Golang

## Build Locally

Change user/pass for nats accordingly in the `main.go` files, then run:

```
cd go-nats-pub && go build .
cd ../go-nats-sub && go build .
```

## Run

```
cd go-nats-sub && ./go-nats-sub
```

In another terminal:
```
cd go-nats-pub && ./go-nats-pub
```

Should see logs from 1st tab from the subscriber.

## Docker Build

Change user/pass for nats accordingly in the `main.go` files, then run:

```
docker build . -t ironcore864/go-nats
```

When running this image, subscriber is started. You can execute the publisher accordingly.
