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

	sub, err := sc.QueueSubscribe("foo", "qfoo", func(m *stan.Msg) {
		if err := m.Ack(); err != nil {
			log.Println(err)
		}
		// fake processing time
		time.Sleep(time.Millisecond * 10)
		fmt.Println(string(m.Data))
	}, stan.MaxInflight(10), stan.AckWait(time.Second), stan.SetManualAckMode())
	if err != nil {
		log.Fatalln(err)
	}
	defer sub.Unsubscribe()

	time.Sleep(1000 * time.Hour)
}
