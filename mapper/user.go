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

func NewUser(db *models.DbConn) *User {
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

func (u *User) UpdateUser(ctx context.Context, user *models.User) error {
	oldUser := &models.User{}
	res := u.DB.First(&oldUser, user.ID)
	if res.Error != nil {
		return res.Error
	}
	oldUser = user
	res = u.DB.Save(&oldUser)
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
