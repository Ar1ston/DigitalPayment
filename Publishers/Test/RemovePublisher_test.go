package Test

import (
	"DigitalPayment/Publishers/Requests"
	"testing"
)

func Test_RemovePublisher_Success(t *testing.T) {
	req := Requests.RequestRemovePublisher{Id: 3}
	_, err := req.Execute()
	if err != nil {
		t.Fail()
	}
}
func Test_RemovePublisher_BadId(t *testing.T) {
	req := Requests.RequestRemovePublisher{Id: 154254}
	_, err := req.Execute()
	if err != nil {
		t.Fail()
	}
}
