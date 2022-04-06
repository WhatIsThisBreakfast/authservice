package handlerutil

import "encoding/json"

type HandlerError struct {
	Message string `json:"error"`
}

func NewError(message string) []byte {
	hError := &HandlerError{
		Message: message,
	}

	createdError, _ := json.Marshal(hError)

	return createdError
}
