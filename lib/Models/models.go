package Models

type Logic interface {
	Decode([]byte) *error
	Validation() []byte
	Execute() ([]byte, *error)
}
