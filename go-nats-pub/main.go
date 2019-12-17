package main

import (
	"fmt"
	"os"

	nats "github.com/nats-io/nats.go"
)

func main() {
	natsConnectionURL := fmt.Sprintf("nat://%s:4222", os.Getenv("NATS_HOST"))

	nc, _ := nats.Connect(natsConnectionURL, nats.UserInfo(os.Getenv("USER"), os.Getenv("PASS")))
	defer nc.Close()

	for i := 1; i <= 10; i++ {
		content := fmt.Sprintf("Hello World %d", i)
		nc.Publish("foo", []byte(content))
		fmt.Println("Published one message! ", content)
	}
	err := nc.Flush()
	if err == nil {
		fmt.Println("Everything has been processed by the server for nc *Conn.")
	}
}
