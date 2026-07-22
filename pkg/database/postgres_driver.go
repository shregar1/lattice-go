package database

import "context"

type PostgresDriver struct {
	ConnectionString string
}

func NewPostgresDriver(connStr string) *PostgresDriver { return &PostgresDriver{ConnectionString: connStr} }
func (p *PostgresDriver) Connect(ctx context.Context) error { return nil }
func (p *PostgresDriver) Disconnect(ctx context.Context) error { return nil }
func (p *PostgresDriver) Query(ctx context.Context, sql string, params ...any) ([]any, error) { return []any{}, nil }
func (p *PostgresDriver) HealthCheck(ctx context.Context) bool { return true }
func (p *PostgresDriver) GetDriverName() string { return "postgres" }
