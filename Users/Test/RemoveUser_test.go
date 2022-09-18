package Test

import (
	"DigitalPayment/Users/Requests"
	"testing"
)

func Test_RemoveUser_Success(t *testing.T) {
	req := Requests.RequestRemoveUser{Id: 4}
	_, err := req.Execute()
	if err != nil {
		t.Fail()
	}
}
func Test_RemoveUser_BadId(t *testing.T) {
	req := Requests.RequestRemoveUser{Id: 1548654}
	_, err := req.Execute()
	if err != nil {
		t.Fail()
	}
}
