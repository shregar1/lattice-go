// Package abstractions provides base interface abstractions for Lattice-Go.
package abstractions

import "context"

type IUnitOfWork interface {
	ExecuteInTransaction(ctx context.Context, work func(txCtx context.Context) error) error
}

type BaseOrchestrator struct {
	UOW IUnitOfWork
}

func (o *BaseOrchestrator) ExecuteInTransaction(ctx context.Context, work func(txCtx context.Context) error, actionName string) error {
	return o.UOW.ExecuteInTransaction(ctx, work)
}
