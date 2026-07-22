package exceptions

func NewConflictException(msg string) *AppException {
	return &AppException{Message: msg, StatusCode: 409, ResponseKey: "CONFLICT"}
}
