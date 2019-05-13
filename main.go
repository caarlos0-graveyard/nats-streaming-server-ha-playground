package main

import (
	"fmt"
	"time"

	"github.com/nats-io/stan"
)

func main() {
	sc, err := stan.Connect("test-cluster", "my-cli")
	if err != nil {
		panic(err)
	}
	defer sc.Close()

	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()

	var i = 1
	for {
		if err := sc.Publish("foo", []byte(fmt.Sprintf("Hello World %d", i))); err != nil {
			panic(err)
		}
		i++
		time.Sleep(time.Millisecond * 100)
	}
}
