package domain

import (
	"strings"
)

type Response struct {
	StatusCode uint32      `json:"status_code"`
	Message    string      `json:"message"`
	Error      interface{} `json:"error"`
	Data       interface{} `json:"data"`
}

func SuccessReponse(statusCode uint32, message string, data ...interface{}) Response {
	return Response{
		StatusCode: statusCode,
		Message:    message,
		Error:      nil,
		Data:       data,
	}
}

func ErrorReponse(statusCode uint32, message string, errorInString string, data interface{}) Response {

	return Response{
		StatusCode: statusCode,
		Message:    message,
		Error:      strings.Split(errorInString, "/n"),
		Data:       data,
	}
}
