package exceptions

func NewForbiddenException(msg string) *AppException {
	return &AppException{Message: msg, StatusCode: 403, ResponseKey: "FORBIDDEN"}
}
