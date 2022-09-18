package main

import (
	"DigitalPayment/Users/Requests"
	"DigitalPayment/Users/lib/db_local"
	"DigitalPayment/lib/NATS"
	"DigitalPayment/lib/logs"
	"DigitalPayment/lib/parameters"
)

func main() {
	serviceName := "Users"

	//Инициализация логов
	logs.Logger.SetService(serviceName).SetRequest("")

	err := parameters.ParamsService.LoadINI("config.ini")
	if err != nil {
		logs.Logger.Errorf("Ошибка чтения ini-файла: %s", err)
		return
	}

	logs.Logger.Infof("HostNATS: %s, PortNATS:%s", parameters.ParamsService.NatsHost, parameters.ParamsService.NatsPort)

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
