package services

import (
	"context"
	"errors"
	"wanpei-backend/mapper"
	"wanpei-backend/models"
	"wanpei-backend/server"
	"wanpei-backend/utils"
)

type User struct {
	UserMapper *mapper.User
	Setting    *server.Settings
}

func NewUser(user *mapper.User, settings *server.Settings) *User {
	return &User{UserMapper: user, Setting: settings}
}

func (u *User) CreateUser(ctx context.Context, user *models.User) error {
	// check for duplicate
	dupUser, err := u.UserMapper.GetUserByEmail(ctx, user.Email)
	if dupUser != nil {
		return errors.New("duplicated user")
	}
	// encrypt password
	user.Password = utils.Sha256WithSalt(user.Password, u.Setting.Sha256Salt)
	// create the user
	err = u.UserMapper.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Login(ctx context.Context, email string, password string) (*models.User, error) {
	password = utils.Sha256WithSalt(password, u.Setting.Sha256Salt)
	user, err := u.UserMapper.GetUserByMap(ctx, map[string]any{"email": email,
		"password": password})
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return user, nil

}
