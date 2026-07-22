package exceptions

func NewConstraintViolationException(msg string) *AppException {
	return &AppException{Message: msg, StatusCode: 400, ResponseKey: "CONSTRAINT_VIOLATION"}
}
