package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
	"net/http"
	"wanpei-backend/models"
	"wanpei-backend/services"
)

// User include all the services that would be used when handling the req
type User struct {
	fx.In
	UserService *services.User
}

func Register(App *gin.Engine, user User) {
	App.POST("/user/register", user.RegisterHandler)
}

func (u *User) RegisterHandler(c *gin.Context) {
	// binding required fields. This will raise an error if empty.
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "empty fields"})
		return
	}
	log.Println(user.Password, user.Email)
	err = u.UserService.CreateUser(c, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Create user failed"})
		return
	}
	c.JSON(200, gin.H{"message": "ok"})
}
