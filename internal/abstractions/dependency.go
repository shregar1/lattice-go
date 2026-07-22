package abstractions

type IDependencyContainer interface {
	Register(token any, instance any)
	Resolve(token any) any
}
