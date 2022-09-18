package Test

import (
	"DigitalPayment/Publishers/Requests"
	"testing"
)

func Test_GetPublisher_Success(t *testing.T) {
	req := Requests.RequestGetPublisher{Id: 3}
	_, err := req.Execute()
	if err != nil {
		t.Fail()
	}
}
func Test_GetPublisher_BadId(t *testing.T) {
	req := Requests.RequestGetPublisher{Id: 21651516}
	_, err := req.Execute()
	if err != nil {
		t.Fail()
	}
}
