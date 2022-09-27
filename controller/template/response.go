package template

import (
	"encoding/json"
	"log"
)

type BaseResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

type BaseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (r BaseResponse[T]) ToJson() []byte {
	b, err := json.Marshal(r)
	if err != nil {
		log.Fatal("error when make response json")
	}
	return b
}

func BaseErrorResponse() *BaseError {
	return &BaseError{
		Code:    -1,
		Message: "default error message",
	}
}

func BaseSuccessResponse() *BaseResponse[any] {
	return &BaseResponse[any]{
		Code:    1,
		Message: "ok",
	}
}
