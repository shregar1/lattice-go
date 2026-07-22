package exceptions

func NewNotFoundException(msg string) *AppException {
	return &AppException{Message: msg, StatusCode: 404, ResponseKey: "NOT_FOUND"}
}
