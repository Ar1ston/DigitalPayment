package Requests

func init() {

}

type RequestChangeBook struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name,omitempty"`
	Genre       string `json:"genre,omitempty"`
	Author      uint64 `json:"author,omitempty"`
	Publisher   uint64 `json:"publisher,omitempty"`
	Description string `json:"description,omitempty"`
}
type ResponseChangeBook struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestChangeBook) Validation() {

}
func (request *RequestChangeBook) Execute() {

}
