package NATS

import (
	"DigitalPayment/lib/crypt"
	"DigitalPayment/lib/logs"
	"DigitalPayment/lib/register_requests"
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
			logs.Logger.Info("Incoming request received")
			in_struct := RequestNats{}
			out_struct := RequestNats{}

			//декодирование входящего сообщения
			err := crypt.Gob_decrypt(m.Data, &in_struct)
			if err != nil {
				logs.Logger.Errorf("ERROR DECRYPT REQUEST: %s", err.Error())
				return
			}
			logs.Logger.Infof("REQUEST FROM: %s TO: %s ReqName: %s", in_struct.From, in_struct.To, in_struct.RequestName)

			//поиск нужной структуры
			findStruct, err := register_requests.FindStruct(in_struct.RequestName)
			if err != nil {
				logs.Logger.Errorf("ERROR FIND ReqName: %s", err)
				return
			}

			//Бизнес логика
			logs.Logger.Info("Start Business logic")
			var rpl []byte
			findStruct.Decode(in_struct.Msg)
			rpl = findStruct.Validation()
			if rpl == nil {
				rpl, _ = findStruct.Execute()
			}
			logs.Logger.Info("End Business logic")

			//заполнение ответа
			out_struct.To = in_struct.From
			out_struct.From = service_name
			out_struct.RequestName = in_struct.RequestName
			out_struct.Msg = rpl

			// Кодирование ответа
			resp_byte, err := crypt.Gob_encrypt(&out_struct)
			if err != nil {
				logs.Logger.Errorf("ERROR ENCRYPT: %s", err.Error())
				return
			}

			//отправка ответа
			err = nats_local.Publish(m.Reply, resp_byte)
			if err != nil {
				logs.Logger.Errorf("ERROR SEND RESPONSE: %s\n", err.Error())
				return
			}
		}()

	})
	defer func() {
		conn.Unsubscribe()
		nats_local.Close()
	}()
	if err != nil {
		logs.Logger.Error("Worker is not running")
		return
	}

	logs.Logger.Infof("Worker %s is running", service_name)

	<-signal_chan

	logs.Logger.Info("Exeting...")
}
