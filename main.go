package main

import (
	"fmt"
	"log"
	"math/rand"
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
	var addr = addrs[rand.Intn(len(addrs))]
	log.Println("connecting to", addr)

	sc, err := stan.Connect("test-cluster", "client-1", stan.Option(func(opts *stan.Options) error {
		opts.NatsURL = addr
		return nil
	}))
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
