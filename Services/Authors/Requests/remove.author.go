package Requests

func init() {

}

type RequestRemoveAuthor struct {
	Id uint64 `json:"id"`
}
type ResponseRemoveAuthor struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestRemoveAuthor) Validation() {

}
func (request *RequestRemoveAuthor) Execute() {

}
