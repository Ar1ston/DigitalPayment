package Requests

type RequestChangePublisher struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
type ResponseChangePublisher struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestChangePublisher) Validation() {

}
func (request *RequestChangePublisher) Execute() {

}
