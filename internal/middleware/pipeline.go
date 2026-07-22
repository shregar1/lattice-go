// Package middleware provides the 15-stage middleware pipeline for Lattice-Go.
package middleware

// PipelineStages lists the 15 stages executed in exact numerical order.
var PipelineStages = []string{
	"1. ExceptionHandlerMiddleware",
	"2. TrustedHostMiddleware",
	"3. SecurityHeadersMiddleware",
	"4. CorsMiddleware",
	"5. CompressionMiddleware",
	"6. RequestTimeoutMiddleware",
	"7. RequestContextMiddleware",
	"8. RequestLoggerMiddleware",
	"9. RateLimitMiddleware",
	"10. AuthenticationMiddleware",
	"11. TenantResolutionMiddleware",
	"12. AuthorizationMiddleware",
	"13. RequestValidationMiddleware",
	"14. ResponseBuilderMiddleware",
	"15. AuditLoggerMiddleware",
}

type MiddlewarePipeline struct{}
