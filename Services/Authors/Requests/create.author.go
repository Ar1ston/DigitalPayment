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

func (request *RequestCreateAuthor) Validation() *error {

	return nil
}
func (request *RequestCreateAuthor) Execute() ([]byte, *error) {

	return nil, nil
}
