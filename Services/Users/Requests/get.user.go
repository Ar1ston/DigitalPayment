package Requests

type RequestGetUser struct {
	Id    uint64 `json:"id,omitempty"`
	Login string `json:"login,omitempty"`
}
type ResponseGetUser struct {
	Id       uint64 `json:"id"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Level    uint64 `json:"level"`
	Errno    uint64 `json:"errno"`
	Error    string `json:"error,omitempty"`
}

func (request *RequestGetUser) Validation() {

}
func (request *RequestGetUser) Execute() {

}
