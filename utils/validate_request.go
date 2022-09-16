package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func ValidateNotNull(ctx *gin.Context, fields []string) error {
	for _, field := range fields {
		if ctx.PostForm(field) == "" {
			return errors.New("empty fields")
		}
	}
	return nil
}
