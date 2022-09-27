package services

import (
	"context"
	"errors"
	"gorm.io/gorm"
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

func (u *User) Login(ctx context.Context, email string, password string) (*models.UserInsensitive, error) {
	password = utils.Sha256WithSalt(password, u.Setting.Sha256Salt)
	user, err := u.UserMapper.GetUserByMap(ctx, map[string]any{"email": email,
		"password": password})
	if err != nil {
		return nil, err
	}
	userInsensitive := utils.ToInsensitiveUser(user)
	if err != nil {
		return nil, err
	}
	return userInsensitive, nil
}

func (u *User) AddGameToUser(ctx context.Context, user *models.User) error {
	//todo: finish the logic here
	return nil
}

func (u *User) ModifyUser(userInsensitive *models.UserInsensitive) error {

	err := u.UserMapper.UpdateUser(context.Background(), userInsensitive)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetUser(ID uint) *models.UserInsensitive {
	user, err := u.UserMapper.GetUserById(context.Background(), ID)

	userInsensitive := &models.UserInsensitive{
		Model: gorm.Model{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
		Username:  user.Username,
		Email:     user.Email,
		Nickname:  user.Nickname,
		Games:     user.Games,
		SteamCode: user.SteamCode,
	}

	if err != nil {
		return nil
	}
	return userInsensitive
}
