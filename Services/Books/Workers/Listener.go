package main

import (
	"DigitalPayment/Services/Authors/Requests"
	"DigitalPayment/lib/NATS"
	"fmt"
)

func main() {

	conn := NATS.ConnectNATS{
		Host: "localhost",
		Port: "4222",
	}
	nats, err := conn.ConnectToNATS()
	if err != nil {
		fmt.Printf("Ошибка подключения к NATS: %s", err)
		return
	}
	Requests.Hello()
	NATS.RunWorker(nats, "Books")
}
