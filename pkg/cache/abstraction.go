package cache
import "context"
type BaseCacheClient interface {
	Get(ctx context.Context, key string) (any, error)
	Set(ctx context.Context, key string, value any, ttlSeconds int) error
	Delete(ctx context.Context, key string) (bool, error)
	GetDriverName() string
}
