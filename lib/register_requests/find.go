package register_requests

import (
	"DigitalPayment/lib/Models"
	"fmt"
)

func FindStruct(method string) (Models.Logic, error) {

	registered, ok := registeredRequests[method]
	if !ok {
		return nil, fmt.Errorf("%s", "Метод не найден")
	}
	msg := registered.CreateRequest()
	if msg == nil {
		return nil, fmt.Errorf("%s", "Ошибка создания структуры запроса")
	}
	return msg, nil
}
