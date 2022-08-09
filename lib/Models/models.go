package Models

type Logic interface {
	Decode([]byte) *error
	Validation() *error
	Execute() ([]byte, *error)
}
