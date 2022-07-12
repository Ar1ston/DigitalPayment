package Requests

type RequestCreatePublisher struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
type ResponseCreatePublisher struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestCreatePublisher) Validation() {

}
func (request *RequestCreatePublisher) Execute() {

}
