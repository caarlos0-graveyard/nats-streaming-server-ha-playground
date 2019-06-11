package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nats-io/stan"
)

func main() {
	sc, err := stan.Connect(
		"test-cluster",
		os.Args[1],
		stan.Pings(1, 3),
		stan.MaxPubAcksInflight(20),
		stan.NatsURL(strings.Join(os.Args[2:], ",")),
		stan.SetConnectionLostHandler(func(con stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v, will retry", reason)
		}),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer sc.Close()

	for {
		if err := sc.Publish("foo", []byte("msg")); err != nil {
			log.Println(err)
		}
		fmt.Print(".")
	}
}
