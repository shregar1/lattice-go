package multitenancy
import "context"
type BaseTenantIsolationStrategy interface {
	ResolveTenantID(ctx context.Context) string
	ApplyIsolationBoundary(ctx context.Context, tenantID string, conn any) any
}
type HeaderBasedTenantIsolationStrategy struct{}
