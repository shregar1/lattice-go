package exceptions

func NewUnauthorizedException(msg string) *AppException {
	return &AppException{Message: msg, StatusCode: 403, ResponseKey: "UNAUTHORIZED"}
}
