package main

import (
	"fmt"

	nats "github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect("nat://nats-test-client:4222", nats.UserInfo("nats_client", "xxx"))
	nc.Publish("foo", []byte("Hello World"))
	fmt.Println("Published one message!")
	nc.Close()
}
