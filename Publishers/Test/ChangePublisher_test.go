package Test

import (
	"DigitalPayment/Publishers/Requests"
	"testing"
)

func Test_ChangePublisher_Success(t *testing.T) {
	req := Requests.RequestChangePublisher{
		Id:          3,
		Name:        "Test_Publisher1",
		Description: "Test_Publisher1",
	}
	_, err := req.Execute()
	if err != nil {
		t.Fatal(err)
	}

}

//
//func Test_ChangePublisher_BadFirstName(t *testing.T) {
//	req := Requests.RequestChangePublisher{
//		Id:          1,
//		Last_name:   "Неизвестен",
//		Description: "Неизвестен",
//	}
//	_, err := req.Execute()
//	if err == nil {
//		t.Fatal()
//	}
//
//}
//func Test_ChangePublisher_BadLastName(t *testing.T) {
//	req := Requests.RequestChangePublisher{
//		Id:          1,
//		First_name:  "Неизвестен",
//		Description: "Неизвестен",
//	}
//	_, err := req.Execute()
//	if err == nil {
//		t.Fatal()
//	}
//}
//func Test_ChangePublisher_BadDesc(t *testing.T) {
//	req := Requests.RequestChangePublisher{
//		Id:         1,
//		First_name: "Неизвестен",
//		Last_name:  "Неизвестен",
//	}
//	_, err := req.Execute()
//	if err == nil {
//		t.Fatal()
//	}
//
//}
//func Test_ChangePublisher_BadId(t *testing.T) {
//	req := Requests.RequestChangePublisher{
//		First_name: "Неизвестен",
//		Last_name:  "Неизвестен",
//	}
//	_, err := req.Execute()
//	if err == nil {
//		t.Fatal()
//	}
//
//}
//func Test_ChangePublisher_NotFoundId(t *testing.T) {
//	req := Requests.RequestChangePublisher{
//		Id:          1245678,
//		First_name:  "Неизвестен",
//		Last_name:   "Неизвестен",
//		Description: " ",
//	}
//	_, err := req.Execute()
//	if err == nil {
//		t.Fatal()
//	}
//
//}
