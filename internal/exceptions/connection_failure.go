package exceptions

func NewConnectionFailureException(msg string) *AppException {
	return &AppException{Message: msg, StatusCode: 503, ResponseKey: "CONNECTION_FAILURE"}
}
