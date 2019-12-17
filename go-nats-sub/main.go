package main

import (
	"fmt"
	"os"
	"runtime"

	nats "github.com/nats-io/nats.go"
)

func main() {
	natsConnectionURL := fmt.Sprintf("nat://%s:4222", os.Getenv("NATS_HOST"))

	nc, _ := nats.Connect(natsConnectionURL, nats.UserInfo(os.Getenv("USER"), os.Getenv("PASS")))
	defer nc.Close()

	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	nc.Flush()
	fmt.Println("Subscribed to one topic!")
	runtime.Goexit()
}
