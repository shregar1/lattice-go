package abstractions

type IOrchestrator interface {
	ExecuteInTransaction(work func() error) error
}
