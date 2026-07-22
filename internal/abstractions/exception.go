package abstractions

type IException interface {
	Error() string
	StatusCode() int
}
