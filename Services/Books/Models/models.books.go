package Models

import "time"

type Books struct {
	Id          uint64    `json:"id"`
	Name        string    `json:"name"`
	Genre       string    `json:"genre"`
	Author      uint64    `json:"author"`
	Publisher   uint64    `json:"publisher"`
	Description string    `json:"description"`
	Added_User  uint64    `json:"addedUser"`
	Added_Time  time.Time `json:"addedTime"`
}
type Logic interface {
	Execute()
	Validation()
}
