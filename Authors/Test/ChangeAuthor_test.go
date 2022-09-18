package Test

import (
	"DigitalPayment/Authors/Requests"
	"testing"
)

func Test_ChangeAuthor_Success(t *testing.T) {
	req := Requests.RequestChangeAuthor{
		Id:          1,
		First_name:  "Неизвестен",
		Last_name:   "Неизвестен",
		Description: "Неизвестен",
	}
	_, err := req.Execute()
	if err != nil {
		t.Fatal(err)
	}

	req.Id = 1
	req.First_name = "Неизвестно"
	req.Last_name = "Неизвестно"
	req.Description = " "
	_, err = req.Execute()
	if err != nil {
		t.Fatal(err)
	}
}

func Test_ChangeAuthor_BadFirstName(t *testing.T) {
	req := Requests.RequestChangeAuthor{
		Id:          1,
		Last_name:   "Неизвестен",
		Description: "Неизвестен",
	}
	_, err := req.Execute()
	if err == nil {
		t.Fatal()
	}

}
func Test_ChangeAuthor_BadLastName(t *testing.T) {
	req := Requests.RequestChangeAuthor{
		Id:          1,
		First_name:  "Неизвестен",
		Description: "Неизвестен",
	}
	_, err := req.Execute()
	if err == nil {
		t.Fatal()
	}
}
func Test_ChangeAuthor_BadDesc(t *testing.T) {
	req := Requests.RequestChangeAuthor{
		Id:         1,
		First_name: "Неизвестен",
		Last_name:  "Неизвестен",
	}
	_, err := req.Execute()
	if err == nil {
		t.Fatal()
	}

}
func Test_ChangeAuthor_BadId(t *testing.T) {
	req := Requests.RequestChangeAuthor{
		First_name: "Неизвестен",
		Last_name:  "Неизвестен",
	}
	_, err := req.Execute()
	if err == nil {
		t.Fatal()
	}

}
func Test_ChangeAuthor_NotFoundId(t *testing.T) {
	req := Requests.RequestChangeAuthor{
		Id:          1245678,
		First_name:  "Неизвестен",
		Last_name:   "Неизвестен",
		Description: " ",
	}
	_, err := req.Execute()
	if err == nil {
		t.Fatal()
	}

}
