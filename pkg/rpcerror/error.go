package rpcerror

import (
	"net/http"
	"strconv"
)

type Error struct {
	Code    int
	Message string
}

func NewError(code int, message string) *Error {
	return &Error{Code: code, Message: message}
}

func NewServerError(message string) *Error {
	return &Error{Code: http.StatusInternalServerError, Message: message}
}

func (e *Error) Error() string {
	return "Error json-rpc: code " + strconv.Itoa(e.Code) + ", message \"" + e.Message + "\""
}
