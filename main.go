package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/nats-io/stan"
)

func main() {
	sc, err := stan.Connect(
		"test-cluster",
		"client-1",
		stan.Pings(1, 3),
		stan.NatsURL(strings.Join(os.Args[1:], ",")),
		stan.SetConnectionLostHandler(func(con stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v, will retry", reason)
		}),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer sc.Close()

	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Print(".")
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer sub.Unsubscribe()

	for {
		if err := sc.Publish("foo", []byte("msg")); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(time.Millisecond * 100)
	}
}
