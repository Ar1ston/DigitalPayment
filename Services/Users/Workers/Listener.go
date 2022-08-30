package main

import (
	"DigitalPayment/Services/Users/Requests"
	"DigitalPayment/lib/NATS"
	"DigitalPayment/lib/logs"
)

func main() {
	serviceName := "Users"

	conn := NATS.ConnectNATS{
		Host: "localhost",
		Port: "4222",
	}
	nats, err := conn.ConnectToNATS()
	if err != nil {
		logs.Logger.Errorf("Ошибка подключения к NATS: %s", err)
		return
	}
	Requests.Hello()
	NATS.RunWorker(nats, serviceName)
}
