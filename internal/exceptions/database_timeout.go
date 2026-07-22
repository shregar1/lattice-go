package exceptions

func NewDatabaseTimeoutException(msg string) *AppException {
	return &AppException{Message: msg, StatusCode: 504, ResponseKey: "DATABASE_TIMEOUT"}
}
