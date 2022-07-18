package Test

import (
	"DigitalPayment/Services/Books/Requests"
	"testing"
)

func Test_ChangeBook_Success(t *testing.T) {
	req := Requests.RequestChangeBook{
		Id:          2,
		Name:        "Неизвестен",
		Genre:       "Неизвестен",
		Author:      2,
		Publisher:   2,
		Description: "Неизвестен",
	}
	_, err := req.Execute()
	if err != nil {
		t.Fatal(err)
	}
}
