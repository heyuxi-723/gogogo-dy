package service

import (
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/models"
)

type UserLoginResponse models.UserLoginResponse

func (q *UserLoginResponse) Register(user *models.User) error {
	err := models.AddUser(user)
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
