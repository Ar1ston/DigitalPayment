package Test

import (
	"DigitalPayment/Books/Requests"
	"testing"
)

func Test_RemoveBook_Success(t *testing.T) {
	req := Requests.RequestRemoveBook{Id: 2}
	_, err := req.Execute()
	if err != nil {
		t.Fail()
	}
}
