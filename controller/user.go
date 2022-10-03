package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
	"net/http"
	"wanpei-backend/controller/template"
	"wanpei-backend/models"
	"wanpei-backend/services"
	"wanpei-backend/utils"
)

// User include all the services that would be used when handling the req
type User struct {
	fx.In
	UserService *services.User
	SessionMgr  *cookie.Store
}

func UserRoutes(App *gin.Engine, user User) {
	// dont need to check login status
	App.POST("/user/register", user.Register)
	App.POST("/user/login", user.Login)
	// check login status routes
	UserGroup := App.Group("/user")
	UserGroup.Use(ValidateLoginStatus)
	UserGroup.POST("/logout", user.LogOut)
	UserGroup.GET("/logout", user.LogOut)
	UserGroup.GET("/current", user.Current)
	UserGroup.POST("/modify", user.Modify)
}

func (u *User) LogOut(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Options(sessions.Options{
		Path:     "/",
		Domain:   "",
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: false,
		SameSite: 0,
	})
	session.Clear()
	if err := session.Save(); err != nil {
		log.Fatal(err)
		return
	}
}

func (u *User) Login(ctx *gin.Context) {
	// use email to login
	type loginForm struct {
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	var params loginForm
	err := ctx.ShouldBind(&params)
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
	session.Options(sessions.Options{
		Path:     "/",
		Domain:   "",
		MaxAge:   60 * 60 * 24, // 24 hours
		Secure:   false,
		HttpOnly: false,
		SameSite: 0,
	})
	session.Set("user", user)
	err = session.Save()
	if err != nil {
		log.Fatal(err)
		return
	}
	ctx.JSON(200, gin.H{"message": "ok"})
}

func (u *User) Current(ctx *gin.Context) {
	session := sessions.Default(ctx)
	if user := session.Get("user"); user != nil {
		userInsensitive := user.(models.UserInsensitive)
		// get latest user info from database
		newUser := u.UserService.GetUser(userInsensitive.ID)
		session.Set("user", newUser)
		_ = session.Save()
		ctx.JSON(200, template.BaseResponse[models.UserInsensitive]{
			Code:    1,
			Message: "success",
			Data:    *newUser,
		})
		return
	} else {
		ctx.JSON(200, template.BaseResponse[any]{
			Code:    -1,
			Message: "no current user",
			Data:    nil,
		})
	}
}

func (u *User) Register(c *gin.Context) {
	// binding required fields. This will raise an error if empty.
	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, template.BaseError{
			Code:    -1,
			Message: "empty fields",
		})
		return
	}
	user := models.User{
		Email:    email,
		Password: password,
	}
	// call user service to handle the logic
	err := u.UserService.CreateUser(c, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Create user failed"})
		return
	}
	c.JSON(http.StatusOK, template.BaseSuccessResponse())
}

func (u *User) Modify(ctx *gin.Context) {
	// validate login
	_, err := utils.ValidateLoginStatus(ctx)
	if err != nil {
		ctx.JSON(400, template.BaseResponse[any]{
			Code:    -1,
			Message: "not logged in",
			Data:    nil,
		})
		return
	}
	// bind to object
	var NewUserInsensitive models.UserInsensitive
	err = ctx.ShouldBindJSON(&NewUserInsensitive)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, template.BaseErrorResponse())
		return
	}
	// send to service layer to handle change user info request
	err = u.UserService.ModifyUser(&NewUserInsensitive)
	if err != nil {
		ctx.JSON(400, template.BaseErrorResponse())
		return
	}
	// return status
	ctx.JSON(200, template.BaseSuccessResponse())
}
