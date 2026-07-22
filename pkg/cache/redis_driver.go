package cache
import "context"
type RedisCacheDriver struct{}
func (r *RedisCacheDriver) Get(ctx context.Context, key string) (any, error) { return nil, nil }
func (r *RedisCacheDriver) Set(ctx context.Context, key string, value any, ttlSeconds int) error { return nil }
func (r *RedisCacheDriver) Delete(ctx context.Context, key string) (bool, error) { return true, nil }
func (r *RedisCacheDriver) GetDriverName() string { return "redis" }
