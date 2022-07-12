package Requests

type RequestGetPublisher struct {
	Id uint64 `json:"id"`
}
type ResponseGetPublisher struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Errno       uint64 `json:"errno"`
	Error       string `json:"error,omitempty"`
}

func (request *RequestGetPublisher) Validation() {

}
func (request *RequestGetPublisher) Execute() {

}
