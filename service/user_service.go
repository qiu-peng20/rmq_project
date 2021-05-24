package service

import (
	"RMQ_Project/DB"
	"RMQ_Project/model"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	IsSuccess(userName string, password string)(user *model.User, ok bool)
	AddUser(user *model.User)(id int64, err error)
}

func NewUserService(userInterface DB.UserInterface) UserServiceInterface {
	return &UserService{userInterface}
}

type UserService struct {
	DB.UserInterface
}

func (u UserService)AddUser(user *model.User)(id int64, err error)  {
	bytePwd, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return id,err
	}
	user.HashPassword = string(bytePwd)
	return u.Insert(user)
}

func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (u UserService)IsSuccess(userName string, password string)(user *model.User, ok bool)  {
	user , err := u.UserInterface.Select(userName)
	if err != nil {
		return
	}
	ok,_ = ValidatePassword(password, user.HashPassword)
	if !ok {
		return &model.User{}, ok
	}
	return 
}

func ValidatePassword(pwd, password string) (ok bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(password), []byte(pwd)); err != nil {
		return false, errors.New("密码不正确")
	}
	return true, nil
}
