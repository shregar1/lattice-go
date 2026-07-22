package cache
import "context"
type InMemoryCacheDriver struct{}
func (i *InMemoryCacheDriver) Get(ctx context.Context, key string) (any, error) { return nil, nil }
func (i *InMemoryCacheDriver) Set(ctx context.Context, key string, value any, ttlSeconds int) error { return nil }
func (i *InMemoryCacheDriver) Delete(ctx context.Context, key string) (bool, error) { return true, nil }
func (i *InMemoryCacheDriver) GetDriverName() string { return "in_memory" }
