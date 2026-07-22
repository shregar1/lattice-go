package migration
import "context"
type BaseMigration interface {
	Up(ctx context.Context) error
	Down(ctx context.Context) error
}
type MigrationRunner struct{}
