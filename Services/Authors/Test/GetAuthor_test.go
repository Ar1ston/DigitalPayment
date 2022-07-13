package Test

import (
	"DigitalPayment/Services/Authors/Requests"
	"testing"
)

func Test_GetAuthor_Success(t *testing.T) {
	req := Requests.RequestGetAuthor{Id: 1}
	_, err := req.Execute()
	if err != nil {
		t.Fail()
	}
}
func Test_GetAuthor_BadId(t *testing.T) {
	req := Requests.RequestGetAuthor{Id: 0}
	err := req.Validation()
	if err == nil {
		t.Fail()
	}
}
