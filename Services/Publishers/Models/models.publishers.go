package Models

type Publisher struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Logic interface {
	Execute()
	Validation()
}
