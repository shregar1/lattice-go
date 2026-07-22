package exceptions

import "fmt"

type AppException struct {
	Message     string `json:"message"`
	StatusCode  int    `json:"statusCode"`
	ResponseKey string `json:"responseKey"`
}

func (e *AppException) Error() string {
	return fmt.Sprintf("[%d %s] %s", e.StatusCode, e.ResponseKey, e.Message)
}

func NewBadRequestException(msg string) *AppException { return &AppException{Message: msg, StatusCode: 400, ResponseKey: "BAD_REQUEST"} }
func NewUnauthenticatedException(msg string) *AppException { return &AppException{Message: msg, StatusCode: 401, ResponseKey: "UNAUTHENTICATED"} }
func NewUnauthorizedException(msg string) *AppException { return &AppException{Message: msg, StatusCode: 403, ResponseKey: "UNAUTHORIZED"} }
func NewForbiddenException(msg string) *AppException { return &AppException{Message: msg, StatusCode: 403, ResponseKey: "FORBIDDEN"} }
func NewNotFoundException(msg string) *AppException { return &AppException{Message: msg, StatusCode: 404, ResponseKey: "NOT_FOUND"} }
func NewConflictException(msg string) *AppException { return &AppException{Message: msg, StatusCode: 409, ResponseKey: "CONFLICT"} }
func NewUnprocessableEntityException(msg string) *AppException { return &AppException{Message: msg, StatusCode: 422, ResponseKey: "UNPROCESSABLE_ENTITY"} }
func NewInternalServerException(msg string) *AppException { return &AppException{Message: msg, StatusCode: 500, ResponseKey: "INTERNAL_ERROR"} }
func NewTimeoutException(msg string) *AppException { return &AppException{Message: msg, StatusCode: 504, ResponseKey: "REQUEST_TIMEOUT"} }
