package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"wanpei-backend/controller/template"
	"wanpei-backend/models"
)

// ValidateLoginStatus is a middleware that validate the login status and store userInsensitive to context
func ValidateLoginStatus(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userObj := session.Get("user")
	if userObj == nil {
		ctx.JSON(http.StatusBadRequest, template.BaseError{
			Code:    -1,
			Message: "not logged in",
		})
		ctx.Abort()
		return
	}
	ctx.Set("user", userObj.(models.UserInsensitive))
}
