package Test

import (
	"DigitalPayment/Services/Publishers/Requests"
	"testing"
)

func Test_GetPublishers_Success(t *testing.T) {
	req := Requests.RequestGetPublishers{}
	_, err := req.Execute()
	if err != nil {
		t.Fail()
	}
}
