package utils

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"wanpei-backend/controller/types"
)

func ValidateLoginStatus(ctx *gin.Context) error {
	session := sessions.Default(ctx)
	user := session.Get("user")
	if user == nil {
		ctx.JSON(404, types.BaseResponse{
			Code:    -1,
			Message: "Not logged in",
			Data:    nil,
		})
		return errors.New("not logged in")
	}
	return nil
}
