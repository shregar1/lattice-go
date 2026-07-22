package exceptions

func NewInternalServerException(msg string) *AppException {
	return &AppException{Message: msg, StatusCode: 500, ResponseKey: "INTERNAL_ERROR"}
}
