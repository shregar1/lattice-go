package abstractions

type IFactory[T any] interface {
	Create(args ...any) T
}
