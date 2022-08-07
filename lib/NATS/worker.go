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
			fmt.Println("ОПОПОПОП ЧТО ТО ПРИШЛО")
			fmt.Printf("ПРИЛЕТ: %s\n", string(m.Data))
			in_struct := RequestNats{}
			out_struct := RequestNats{}

			fmt.Println("Декодирование 1")
			//декодирование входящего сообщения

			//TODO мб тут пиздень
			var buff = bytes.NewBuffer(m.Data)
			fmt.Println("Запись удалась")
			dec := gob.NewDecoder(buff)
			fmt.Println("Декодер создан")
			err := dec.Decode(&in_struct)
			if err != nil {
				fmt.Printf("Декодирование пошло по пизде: %s\n", err)
				return
			}
			fmt.Println("Декодирование 1. Конец")

			fmt.Printf("FROM: %s TO: %s ReqName: %s\n", in_struct.From, in_struct.To, in_struct.RequestName)
			fmt.Println("Поиск структуры")
			//поиск нужной структуры
			findStruct, err := register_requests.FindStruct(in_struct.RequestName)
			if err != nil {
				return
			}
			fmt.Println("Поиск структуры. Конец")

			fmt.Println("Бизнес логика")
			//Бизнес логика
			findStruct.Validation()
			rpl, _ := findStruct.Execute()

			fmt.Println("Бизнес логика. Конец.")

			//заполнение ответа
			out_struct.To = in_struct.From
			out_struct.From = service_name
			out_struct.RequestName = in_struct.RequestName
			out_struct.Msg = rpl

			fmt.Println("Кодирование ответа")
			// Кодирование ответа
			var buff2 bytes.Buffer
			enc := gob.NewEncoder(&buff2)
			err = enc.Encode(out_struct)
			if err != nil {
				return
			}
			fmt.Println("Кодирование ответа. Конец.")

			fmt.Println("Отправка ответа.")
			//отправка ответа
			err = nats_local.Publish(m.Reply, buff2.Bytes())
			if err != nil {
				return
			}
			fmt.Println("Отправка ответа. Конец.")
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
