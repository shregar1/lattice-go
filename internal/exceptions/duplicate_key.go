package exceptions

func NewDuplicateKeyException(msg string) *AppException {
	return &AppException{Message: msg, StatusCode: 409, ResponseKey: "DUPLICATE_KEY"}
}
