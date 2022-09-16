package services

import (
	"context"
	"errors"
	"wanpei-backend/mapper"
	"wanpei-backend/models"
)

type User struct {
	UserMapper *mapper.User
}

func NewUser(user *mapper.User) *User {
	return &User{UserMapper: user}
}

func (u *User) CreateUser(ctx context.Context, user *models.User) error {
	user, err := u.UserMapper.GetUserByEmail(ctx, user.Email)
	if user != nil {
		return errors.New("duplicated user")
	}
	err = u.UserMapper.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
