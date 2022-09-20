package utils

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"wanpei-backend/models"
)

func ValidateLoginStatus(ctx *gin.Context) (*models.User, error) {
	session := sessions.Default(ctx)
	user := session.Get("user")
	if user == nil {
		return nil, errors.New("not logged in")
	}
	userObj := user.(models.User)
	return &userObj, nil
}
