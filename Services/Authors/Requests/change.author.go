package Requests

func init() {

}

type RequestChangeAuthor struct {
	Id          uint64 `json:"id"`
	First_name  string `json:"firstName,omitempty"`
	Last_name   string `json:"lastName,omitempty"`
	Description string `json:"description,omitempty"`
}
type ResponseChangeAuthor struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestChangeAuthor) Validation() *error {

	return nil
}
func (request *RequestChangeAuthor) Execute() ([]byte, *error) {

	return nil, nil
}
