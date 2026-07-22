package exceptions

func NewConcurrencyConflictException(msg string) *AppException {
	return &AppException{Message: msg, StatusCode: 409, ResponseKey: "CONCURRENCY_CONFLICT"}
}
