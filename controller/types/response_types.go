package types

import (
	"encoding/json"
	"log"
)

type BaseResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func (r BaseResponse) ToJson() []byte {
	b, err := json.Marshal(r)
	if err != nil {
		log.Fatal("error when make response json")
	}
	return b
}

func BaseErrorResponse() *BaseResponse {
	return &BaseResponse{
		Code:    -1,
		Message: "default error message",
		Data:    nil,
	}
}
