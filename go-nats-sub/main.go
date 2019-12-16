package main

import (
	"fmt"
	"runtime"

	nats "github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect("nat://nats-test-client:4222", nats.UserInfo("nats_client", "xxx"))

	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	fmt.Println("Subscribed to one topic!")
	runtime.Goexit()
	nc.Close()
}
