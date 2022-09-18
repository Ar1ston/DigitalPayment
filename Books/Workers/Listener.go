package main

import (
	"DigitalPayment/Books/Requests"
	"DigitalPayment/Books/lib/db_local"
	"DigitalPayment/lib/NATS"
	"DigitalPayment/lib/logs"
	"DigitalPayment/lib/parameters"
)

func main() {
	serviceName := "Books"

	//Инициализация логов
	logs.Logger.SetService(serviceName).SetRequest("")

	err := parameters.ParamsService.LoadINI("config.ini")
	if err != nil {
		logs.Logger.Errorf("Ошибка чтения ini-файла: %s", err)
		return
	}

	conn := NATS.ConnectNATS{
		Host: parameters.ParamsService.NatsHost,
		Port: parameters.ParamsService.NatsPort,
	}
	nats, err := conn.ConnectToNATS()
	if err != nil {
		logs.Logger.Errorf("Ошибка подключения к NATS: %s", err)
		return
	}

	err = db_local.InitDB()
	if err != nil {
		logs.Logger.Errorf("Ошибка подключения к базе данных: %s", err)
		return
	}

	Requests.Hello()
	NATS.RunWorker(nats, serviceName)
}
