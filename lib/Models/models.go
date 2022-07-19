package Models

type Logic interface {
	Validation() *error
	Execute() ([]byte, *error)
}
