package database

import "context"

type SqliteDriver struct {
	FilePath string
}

func NewSqliteDriver(path string) *SqliteDriver { return &SqliteDriver{FilePath: path} }
func (s *SqliteDriver) Connect(ctx context.Context) error { return nil }
func (s *SqliteDriver) Disconnect(ctx context.Context) error { return nil }
func (s *SqliteDriver) Query(ctx context.Context, sql string, params ...any) ([]any, error) { return []any{}, nil }
func (s *SqliteDriver) HealthCheck(ctx context.Context) bool { return true }
func (s *SqliteDriver) GetDriverName() string { return "sqlite" }
