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
		os.Args[1],
		stan.Pings(1, 30),
		stan.MaxPubAcksInflight(20),
		stan.PubAckWait(5*time.Second),
		stan.NatsURL(strings.Join(os.Args[2:], ",")),
		stan.SetConnectionLostHandler(func(con stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v, will retry", reason)
		}),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer sc.Close()

	var i int64
	for {
		if err := sc.Publish("foo", []byte(fmt.Sprintf("msg %d", i))); err != nil {
			log.Println(err)
		}
		i++
		fmt.Print(".")
	}
}
