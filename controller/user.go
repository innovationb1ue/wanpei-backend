package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
	"wanpei-backend/models"
	"wanpei-backend/services"
)

// User include all the services that would be used when handling the req
type User struct {
	fx.In
	UserService *services.User
	SessionMgr  *cookie.Store
}

func UserRoutes(App *gin.Engine, user User) {
	App.POST("/user/register", user.RegisterHandler)
	App.POST("/user/login", user.Login)
	App.GET("/user/current", user.CurrentUser)
}

func (u *User) Login(ctx *gin.Context) {
	// use email to login
	type loginForm struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var params loginForm
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "empty fields"})
		return
	}
	user, err := u.UserService.Login(ctx, params.Email, params.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "user not exist or wrong password"})
		return
	}
	session := sessions.Default(ctx)
	session.Set("user", user)
	_ = session.Save()
	ctx.JSON(200, gin.H{"message": "ok"})
}

func (u *User) CurrentUser(ctx *gin.Context) {
	session := sessions.Default(ctx)
	if user := session.Get("user"); user != nil {
		ctx.JSON(200, gin.H{"message": "ok", "data": user})
		return
	} else {
		ctx.JSON(500, gin.H{"message": "no current user"})
	}
}

func (u *User) RegisterHandler(c *gin.Context) {
	// binding required fields. This will raise an error if empty.
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "empty fields"})
		return
	}
	// call services to handle the logic
	err = u.UserService.CreateUser(c, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Create user failed"})
		return
	}
	c.JSON(200, gin.H{"message": "ok"})
}
