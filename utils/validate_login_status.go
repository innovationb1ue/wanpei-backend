package utils

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"wanpei-backend/models"
)

func ValidateLoginStatus(ctx *gin.Context) (*models.UserInsensitive, error) {
	session := sessions.Default(ctx)
	user := session.Get("user")
	userObj, ok := user.(models.UserInsensitive)

	if user == nil || !ok || userObj.ID < 0 {
		return nil, errors.New("not logged in")
	}
	return &userObj, nil
}
