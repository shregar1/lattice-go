package exceptions

func NewUnauthenticatedException(msg string) *AppException {
	return &AppException{Message: msg, StatusCode: 401, ResponseKey: "UNAUTHENTICATED"}
}
