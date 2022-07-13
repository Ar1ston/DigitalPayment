package Test

import (
	"DigitalPayment/Services/Authors/Requests"
	"testing"
)

func Test_GetAuthors_Success(t *testing.T) {
	req := Requests.RequestGetAuthors{}
	_, err := req.Execute()
	if err != nil {
		t.Fail()
	}
}
