package abstractions

type IMiddleware interface {
	Handle(next any) any
}
