package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/models"
	"github.com/RaymondCode/simple-demo/utils"
)

type UserLoginResponse models.UserLoginResponse

func (q *UserLoginResponse) Register(username string, password string) error {
	err := utils.ValidateRegister(username, password, "register")
	if err != nil {
		return err
	}
	newPassword, err := utils.BcryptMake([]byte(password))
	if err != nil {
		return err
	}

	user := &models.User{
		Name:     username,
		Password: newPassword,
	}

	err = models.AddUser(user)
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

func (q *UserLoginResponse) Login(username string, password string) error {
	err := utils.ValidateRegister(username, password, "login")
	if err != nil {
		return err
	}

	user := models.QueryUserLogin(username)

	if isChecked := utils.BcryptMakeCheck([]byte(password), user.Password); !isChecked {
		return errors.New("密码错误，请重试")
	}

	//颁发token
	token, err := middleware.ReleaseToken(models.User{
		Id:   user.Id,
		Name: username,
	})
	if err != nil {
		return err
	}
	q.Token = token
	q.UserId = user.Id
	return nil
}
