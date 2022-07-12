package Models

type Authors struct {
	Id          uint64 `json:"id"`
	First_name  string `json:"firstName"`
	Last_name   string `json:"lastName"`
	Description string `json:"description"`
}
type Logic interface {
	Validation()
	Execute()
}
