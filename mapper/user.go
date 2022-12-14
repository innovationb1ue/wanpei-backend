package mapper

import (
	"context"
	"gorm.io/gorm"
	"log"
	"wanpei-backend/models"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *DbConn) *User {
	return &User{DB: db.Conn}
}

func (u *User) CreateUser(ctx context.Context, user *models.User) error {
	res := u.DB.Create(&user)
	if res.Error != nil {
		log.Fatal("Create user failed")
		return res.Error
	}
	return nil
}

func (u *User) DeleteUserById(ctx context.Context, userId string) error {
	res := u.DB.Delete(&models.User{}, userId)
	if res.Error != nil {
		log.Fatal("Delete user failed")
		return res.Error
	}
	return nil
}

func (u *User) UpdateUser(ctx context.Context, user *models.UserInsensitive) error {

	res := u.DB.Model(&user).Updates(user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *User) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	res := u.DB.Where("email = ? ", email).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (u *User) GetUserByMap(ctx context.Context, conditions map[string]any) (*models.User, error) {
	var user models.User
	res := u.DB.Where(conditions).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (u *User) GetUserById(ctx context.Context, Id uint) (*models.User, error) {
	var user models.User
	res := u.DB.First(&user, Id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (u *User) AddGameById(ctx context.Context, games []string, userId string) error {
	// a valid userId should be passed in. Do checks in service layer.
	var user models.User
	// get by primary key ID
	u.DB.First(&user, userId)
	// todo: finish the logic here
	return nil
}
