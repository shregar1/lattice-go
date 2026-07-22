package database

import "context"

type BaseDatabaseDriver interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Query(ctx context.Context, sql string, params ...any) ([]any, error)
	HealthCheck(ctx context.Context) bool
	GetDriverName() string
}
