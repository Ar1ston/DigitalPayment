package Test

import (
	"DigitalPayment/Users/Requests"
	"testing"
)

func Test_CreateUser_Success(t *testing.T) {
	req := Requests.RequestCreateUser{
		Login:    "User_Test",
		Password: "Password_Tests",
		Name:     "Name_Test",
	}
	_, err := req.Execute()
	if err != nil {
		t.Fail()
	}
}
func Test_CreateUser_BadUser(t *testing.T) {
	req := Requests.RequestCreateUser{
		Login:    "Неизвестно",
		Password: "",
		Name:     "Неизвестно",
	}
	_, err := req.Execute()
	if err != nil {
		t.Fail()
	}
}
