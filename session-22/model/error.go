package model

import "fmt"

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

var (
	UnexpectedError     = ErrorResponse{Message: "UNEXPECTED_EXCEPTION", Code: 500}
	NotFoundError       = ErrorResponse{Message: "NOT_FOUND", Code: 404}
	InValidRequestError = ErrorResponse{Message: "INVALID_BODY", Code: 403}
	UnAuthorized        = ErrorResponse{Message: "UNAUTHORIZED", Code: 401}
	InvalidCredential   = ErrorResponse{Message: "WRONG_CREDENTIAL", Code: 403}
)
