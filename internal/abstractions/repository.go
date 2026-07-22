package abstractions

import "context"

type IBaseRepository[TAny any, TId any] interface {
	FindByID(ctx context.Context, id TId) (*TAny, error)
	Create(ctx context.Context, entity *TAny) (*TAny, error)
}
