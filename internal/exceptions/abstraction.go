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
