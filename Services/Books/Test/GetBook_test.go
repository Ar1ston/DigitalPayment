package Test

import (
	"DigitalPayment/Services/Books/Requests"
	"testing"
)

func Test_GetBook_Success(t *testing.T) {
	req := Requests.RequestGetBook{Id: 2}
	_, err := req.Execute()
	if err != nil {
		t.Fail()
	}
}
