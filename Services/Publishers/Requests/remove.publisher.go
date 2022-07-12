package Requests

type RequestRemovePublisher struct {
	Id uint64 `json:"id"`
}
type ResponseRemovePublisher struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestRemovePublisher) Validation() {

}
func (request *RequestRemovePublisher) Execute() {

}
