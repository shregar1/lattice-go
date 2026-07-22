package exceptions

func NewUnprocessableEntityException(msg string) *AppException {
	return &AppException{Message: msg, StatusCode: 422, ResponseKey: "UNPROCESSABLE_ENTITY"}
}
