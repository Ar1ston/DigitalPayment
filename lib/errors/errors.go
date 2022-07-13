package errors

type Error int64

const (
	INTERNAL_ERROR Error = 500
	NOT_FOUND      Error = 400
)

func (err Error) toString() string {
	switch err {
	case INTERNAL_ERROR:
		return "Internal server error"
	case NOT_FOUND:
		return "Object not found"

	}
	return "Unknown error"
}
