package service

import (
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/models"
)

type UserLoginResponse models.UserLoginResponse

func (q *UserLoginResponse) Register(user *models.User) error {
	err := AddUser(user)
	if err != nil {
		return err
	}

	//颁发token
	token, err := middleware.ReleaseToken(*user)
	if err != nil {
		return err
	}
	q.Token = token
	q.UserId = user.Id
	return nil
}

func AddUser(user *models.User) error {
	return config.Config.DB.Create(&user).Error
}
func IsUserExistByUsername(username string) bool {
	var user models.User
	config.Config.DB.Where("name = ?", username).First(&user)
	if user.Id == 0 {
		return false
	}
	return true
}
