package Requests

type RequestCreateUser struct {
	Id       uint64 `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Level    uint64 `json:"level"`
}
type ResponseCreateUser struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestCreateUser) Validation() {

}
func (request *RequestCreateUser) Execute() {

}
