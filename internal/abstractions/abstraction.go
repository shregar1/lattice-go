// Package abstractions provides base interface abstractions for Lattice-Go.
package abstractions

import "context"

// IModel represents a base persistent entity model.
type IModel interface {
	GetID() string
	GetURN() string
}

// IRepository represents base persistence operations.
type IRepository[T any, ID any] interface {
	FindByID(ctx context.Context, id ID) (*T, error)
	Create(ctx context.Context, entity *T) (*T, error)
}

// IService represents a base single-capability service.
type IService interface {
	ServiceName() string
}

// IOrchestrator represents a base workflow orchestrator.
type IOrchestrator interface {
	OrchestratorName() string
}
