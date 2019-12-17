package main

import (
	"fmt"
	"os"

	nats "github.com/nats-io/nats.go"
)

func main() {
	natsConnectionURL := fmt.Sprintf("nat://%s:4222", os.Getenv("NATS_HOST"))
	nc, _ := nats.Connect(natsConnectionURL, nats.UserInfo(os.Getenv("USER"), os.Getenv("PASS")))
	nc.Publish("foo", []byte("Hello World"))
	fmt.Println("Published one message!")
	nc.Close()
}
