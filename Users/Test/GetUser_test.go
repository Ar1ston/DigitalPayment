package Test

import (
	"DigitalPayment/Users/Requests"
	"testing"
)

func Test_GetUser_Success(t *testing.T) {
	req := Requests.RequestGetUser{Id: 1}
	_, err := req.Execute()
	if err != nil {
		t.Fail()
	}
}
func Test_GetUser_BadUser(t *testing.T) {
	req := Requests.RequestGetUser{Id: 156454}
	_, err := req.Execute()
	if err != nil {
		t.Fail()
	}
}
