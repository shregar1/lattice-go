// Package abstractions provides base interface abstractions for Lattice-Go.
package abstractions

import "context"

type IQueryCriteria map[string]any

type IQueryResult[T any] struct {
	Items      []T `json:"items"`
	Total      int `json:"total"`
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"totalPages"`
}

type IBaseRepository[T any, TId any] interface {
	FindByID(ctx context.Context, id TId) (*T, error)
	FindByURN(ctx context.Context, urn string) (*T, error)
	FindOne(ctx context.Context, criteria IQueryCriteria) (*T, error)
	FindAll(ctx context.Context, criteria IQueryCriteria) ([]T, error)
	FindPaginated(ctx context.Context, criteria IQueryCriteria) (*IQueryResult[T], error)
	Create(ctx context.Context, entity *T) (*T, error)
	CreateMany(ctx context.Context, entities []T) ([]T, error)
	Update(ctx context.Context, id TId, entity map[string]any) (*T, error)
	UpdateMany(ctx context.Context, criteria IQueryCriteria, data map[string]any) (int64, error)
	Delete(ctx context.Context, id TId) (bool, error)
	SoftDelete(ctx context.Context, id TId) (bool, error)
	Restore(ctx context.Context, id TId) (bool, error)
	DeleteMany(ctx context.Context, criteria IQueryCriteria) (int64, error)
	Count(ctx context.Context, criteria IQueryCriteria) (int64, error)
	Exists(ctx context.Context, criteria IQueryCriteria) (bool, error)
}
