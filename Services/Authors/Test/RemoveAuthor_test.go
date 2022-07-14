package Test

import (
	"DigitalPayment/Services/Authors/Requests"
	"testing"
)

func Test_RemoveAuthor_Success(t *testing.T) {
	req := Requests.RequestRemoveAuthor{Id: 1}
	_, err := req.Execute()
	if err != nil {
		t.Fail()
	}
}
func Test_RemoveAuthor_BadId(t *testing.T) {
	req := Requests.RequestRemoveAuthor{}
	err := req.Validation()
	if err == nil {
		t.Fail()
	}
}
