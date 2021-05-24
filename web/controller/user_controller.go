package controller

import (
	"RMQ_Project/model"
	"RMQ_Project/service"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type UserController struct {
	Service service.UserServiceInterface
	Session *sessions.Session
}

func (u *UserController)PostRegister(ctx iris.Context)  {
	var (
		userName = ctx.FormValue("userName")
		HashPassword = ctx.FormValue("HashPassword")
	)
	user := &model.User{
		UserName:userName,
		HashPassword:HashPassword,
	}
	_, err := u.Service.AddUser(user)
	if err != nil {
		fmt.Print(err)
	}
	_, _ = ctx.WriteString("success")
}


