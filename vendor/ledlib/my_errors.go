package ledlib

type ErrorNoObject struct {
}

func (e ErrorNoObject) Error() string {
	return "object is empty"
}
