package Requests

func init() {

}

type RequestCreateAuthor struct {
	Id          uint64 `json:"id"`
	First_name  string `json:"firstName,omitempty"`
	Last_name   string `json:"lastName,omitempty"`
	Description string `json:"description,omitempty"`
}
type ResponseCreateAuthor struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestCreateAuthor) Validation() {

}
func (request *RequestCreateAuthor) Execute() {

}
