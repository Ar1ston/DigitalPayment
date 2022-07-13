package Requests

func init() {

}

type RequestGetAuthors struct{}
type ResponseGetAuthors struct {
	Id          uint64 `json:"id"`
	First_name  string `json:"firstName,omitempty"`
	Last_name   string `json:"lastName,omitempty"`
	Description string `json:"description,omitempty"`
	Errno       uint64 `json:"errno"`
	Error       string `json:"error,omitempty"`
}

func (request *RequestGetAuthors) Validation() *error {

	return nil
}
func (request *RequestGetAuthors) Execute() ([]byte, *error) {

	return nil, nil
}
