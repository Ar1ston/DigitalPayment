package Requests

func init() {

}

type RequestGetAuthor struct {
	Id uint64 `json:"id"`
}
type ResponseGetAuthor struct {
	Id          uint64 `json:"id"`
	First_name  string `json:"firstName,omitempty"`
	Last_name   string `json:"lastName,omitempty"`
	Description string `json:"description,omitempty"`
	Errno       uint64 `json:"errno"`
	Error       string `json:"error,omitempty"`
}

func (request *RequestGetAuthor) Validation() {

}
func (request *RequestGetAuthor) Execute() {

}
