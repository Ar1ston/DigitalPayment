package Test

import (
	"DigitalPayment/Services/Books/Requests"
	"testing"
)

func Test_GetBooks_Success(t *testing.T) {
	req := Requests.RequestGetBooks{}
	_, err := req.Execute()
	if err != nil {
		t.Fail()
	}
}
