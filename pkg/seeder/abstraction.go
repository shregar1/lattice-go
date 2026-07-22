package seeder
import "context"
type BaseSeeder interface {
	Seed(ctx context.Context) error
}
type SeederRunner struct{}
