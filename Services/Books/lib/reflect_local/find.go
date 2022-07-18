package reflect_local

import (
	"DigitalPayment/Services/Books/Models"
	"fmt"
	"net/http"
)

func FindStruct(r *http.Request) (Models.Logic, error) {

	registered, ok := registeredMessages[r.Header.Get("method")]
	if !ok {
		return nil, fmt.Errorf("%s", "Метод не найден")
	}
	msg := registered.CreateRequest()
	if msg == nil {
		return nil, fmt.Errorf("%s", "Ошибка создания структуры запроса")
	}
	return msg, nil
}
