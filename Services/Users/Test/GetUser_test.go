package Test

import (
	"DigitalPayment/Services/Users/Requests"
	"testing"
)

func Test_GetUser_Success(t *testing.T) {
	req := Requests.RequestGetUser{Login: "User_Test1"}
	_, err := req.Execute()
	if err != nil {
		t.Fail()
	}
}
func Test_GetUser_BadUser(t *testing.T) {
	req := Requests.RequestGetUser{Login: "JAGBOGWBAIG"}
	_, err := req.Execute()
	if err != nil {
		t.Fail()
	}
}
