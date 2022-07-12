package Requests

type RequestRemoveUser struct {
	Id uint64 `json:"id"`
}
type ResponseRemoveUser struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestRemoveUser) Validation() {

}
func (request *RequestRemoveUser) Execute() {

}
