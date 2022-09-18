package Test

import (
	"DigitalPayment/Publishers/Requests"
	"testing"
)

func Test_CreatePublisher_Success(t *testing.T) {
	req := Requests.RequestCreatePublisher{
		Name:        "Test_Publisher",
		Description: "Test_Publisher",
	}
	_, err := req.Execute()
	if err != nil {
		t.Fail()
	}
}
