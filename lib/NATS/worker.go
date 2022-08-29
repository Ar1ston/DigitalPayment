package NATS

import (
	"DigitalPayment/lib/crypt"
	"DigitalPayment/lib/register_requests"
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
			fmt.Println("Incoming request received")
			in_struct := RequestNats{}
			out_struct := RequestNats{}

			//декодирование входящего сообщения
			err := crypt.Gob_decrypt(m.Data, &in_struct)
			if err != nil {
				fmt.Printf("ERROR DECRYPT REQUEST: %s\n", err.Error())
				return
			}
			fmt.Printf("REQUEST FROM: %s TO: %s ReqName: %s\n", in_struct.From, in_struct.To, in_struct.RequestName)

			//поиск нужной структуры
			findStruct, err := register_requests.FindStruct(in_struct.RequestName)
			if err != nil {
				fmt.Printf("ERROR FIND ReqName: %s\n", err)
				return
			}

			//Бизнес логика
			fmt.Printf("%s\n", "Start Business logic")
			var rpl []byte
			findStruct.Decode(in_struct.Msg)
			rpl = findStruct.Validation()
			if rpl == nil {
				rpl, _ = findStruct.Execute()
			}
			fmt.Printf("%s\n", "End Business logic")

			//заполнение ответа
			out_struct.To = in_struct.From
			out_struct.From = service_name
			out_struct.RequestName = in_struct.RequestName
			out_struct.Msg = rpl

			// Кодирование ответа
			resp_byte, err := crypt.Gob_encrypt(&out_struct)
			if err != nil {
				fmt.Printf("ERROR ENCRYPT: %s\n", err.Error())
				return
			}

			//отправка ответа
			err = nats_local.Publish(m.Reply, resp_byte)
			if err != nil {
				fmt.Printf("ERROR SEND RESPONSE: %s\n", err.Error())
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
