package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/nats-io/stan"
)

var addrs = []string{
	"nats://localhost:4221",
	"nats://localhost:4222",
	"nats://localhost:4223",
}

func main() {
	for {
		if err := run(); err != nil {
			log.Println(err)
		}
	}
}

func run() error {
	sc, err := stan.Connect("test-cluster", "client-1", stan.NatsURL(strings.Join(addrs, ",")))
	if err != nil {
		return err
	}
	defer sc.Close()

	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Print(".")
	})
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		if err := sc.Publish("foo", []byte("msg")); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 100)
	}
}
