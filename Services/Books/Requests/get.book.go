package Requests

import "time"

type RequestGetBook struct {
	Id   uint64 `json:"id"`
	Name string `json:"name,omitempty"`
}
type ResponseGetBook struct {
	Id          uint64    `json:"id"`
	Name        string    `json:"name,omitempty"`
	Genre       string    `json:"genre,omitempty"`
	Author      uint64    `json:"author,omitempty"`
	Publisher   uint64    `json:"publisher,omitempty"`
	Added_User  uint64    `json:"addedUser,omitempty"`
	Added_Time  time.Time `json:"addedTime,omitempty"`
	Description string    `json:"description,omitempty"`
	Errno       uint64    `json:"errno"`
	Error       string    `json:"error,omitempty"`
}

func (request *RequestGetBook) Validation() {

}
func (request *RequestGetBook) Execute() {

}
