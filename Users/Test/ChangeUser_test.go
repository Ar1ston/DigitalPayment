package Test

import (
	"DigitalPayment/Users/Requests"
	"testing"
)

func Test_ChangeUser_Success(t *testing.T) {
	req := Requests.RequestChangeUser{
		Id:       4,
		Login:    "User_Test1",
		Password: "User_Test1",
		Name:     "User_Test1",
	}
	_, err := req.Execute()
	if err != nil {
		t.Fatal(err)
	}
}

func Test_ChangeUser_NotFoundId(t *testing.T) {
	req := Requests.RequestChangeUser{
		Id:       1245678,
		Login:    "Неизвестен",
		Password: "Неизвестен",
		Name:     " ",
	}
	_, err := req.Execute()
	if err != nil {
		t.Fatal()
	}

}
