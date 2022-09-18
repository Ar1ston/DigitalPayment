package Test

import (
	"DigitalPayment/Authors/Requests"
	"testing"
)

func Test_CreateAuthor_Success(t *testing.T) {
	req := Requests.RequestCreateAuthor{
		First_name:  "Николай",
		Last_name:   "Толстой",
		Description: "Великий русский писатель",
	}
	_, err := req.Execute()
	if err != nil {
		t.Fail()
	}
}
func Test_CreateAuthor_BadFirstName(t *testing.T) {
	req := Requests.RequestCreateAuthor{
		Last_name:   "Толстой",
		Description: "Великий русский писатель",
	}
	_, err := req.Execute()
	if err == nil {
		t.Fail()
	}
}
func Test_CreateAuthor_BadLastName(t *testing.T) {
	req := Requests.RequestCreateAuthor{
		First_name:  "Николай",
		Description: "Великий русский писатель",
	}
	_, err := req.Execute()
	if err == nil {
		t.Fail()
	}
}
func Test_CreateAuthor_BadAuthor(t *testing.T) {
	req := Requests.RequestCreateAuthor{
		First_name:  "Неизвестно",
		Last_name:   "Неизвестно",
		Description: "",
	}
	_, err := req.Execute()
	if err == nil {
		t.Fail()
	}
}
