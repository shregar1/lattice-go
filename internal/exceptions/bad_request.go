package exceptions

func NewBadRequestException(msg string) *AppException {
	return &AppException{Message: msg, StatusCode: 400, ResponseKey: "BAD_REQUEST"}
}
