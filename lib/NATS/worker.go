package NATS

import (
	"DigitalPayment/lib/register_requests"
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/nats-io/nats.go"
	"os"
	"os/signal"
	"syscall"
)

func RunWorker(nats_local *nats.Conn, service_name string) {

	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	conn, err := nats_local.QueueSubscribe(service_name, service_name, func(m *nats.Msg) {

		go func() {
			req := RequestNats{}
			var rplBytes bytes.Buffer
			dec := gob.NewDecoder(&rplBytes)
			err := dec.Decode(&req)
			if err != nil {
				return
			}

			findStruct, err := register_requests.FindStruct(req.RequestName)
			if err != nil {
				return
			}

			findStruct.Validation()
			rpl, _ := findStruct.Execute()
			req.Msg = rpl

			err = nats_local.Publish(m.Reply, toBytes(req))
			if err != nil {
				return
			}
		}()

	})
	defer func() {
		conn.Unsubscribe()
		nats_local.Close()
	}()
	if err != nil {
		fmt.Printf("Worker is not running\n")
		return
	}

	fmt.Printf("Worker %s is running\n", service_name)

	<-signal_chan

	fmt.Printf("Exeting...\n")
}
