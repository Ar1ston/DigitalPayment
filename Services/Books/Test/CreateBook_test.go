package Test

import (
	"DigitalPayment/Services/Books/Requests"
	"testing"
)

func Test_CreateBook_Success(t *testing.T) {
	req := Requests.RequestCreateBook{
		Name:        "Test_Book",
		Genre:       "Test_Book",
		Author:      1,
		Publisher:   1,
		Added_User:  1,
		Description: "Test_Book",
	}
	_, err := req.Execute()
	if err != nil {
		t.Fail()
	}
}
