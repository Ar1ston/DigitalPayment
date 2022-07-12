package Models

type User struct {
	Id       uint64 `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Level    uint64 `json:"level"`
}
type Logic interface {
	Validation()
	Execute()
}
