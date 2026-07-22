package exceptions

func NewTimeoutException(msg string) *AppException {
	return &AppException{Message: msg, StatusCode: 504, ResponseKey: "REQUEST_TIMEOUT"}
}
