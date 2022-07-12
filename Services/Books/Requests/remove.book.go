package Requests

import "time"

type RequestRemoveBook struct {
	Name        string    `json:"name,omitempty"`
	Genre       string    `json:"genre,omitempty"`
	Author      uint64    `json:"author,omitempty"`
	Publisher   uint64    `json:"publisher,omitempty"`
	Added_User  uint64    `json:"addedUser,omitempty"`
	Added_Time  time.Time `json:"addedTime,omitempty"`
	Description string    `json:"description,omitempty"`
}
type ResponseRemoveBook struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestRemoveBook) Validation() {

}
func (request *RequestRemoveBook) Execute() {

}
