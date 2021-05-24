package controller

import (
	"RMQ_Project/model"
	"RMQ_Project/service"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

type UserController struct {
	Service service.UserServiceInterface
	Session *sessions.Session
}

type Success struct {
	Message string `json:"message"`
	Status string `json:"status"`
}

func (u *UserController) PostRegister(Ctx iris.Context) {
	var data model.User
	var (
		err = Ctx.ReadJSON(&data)
	)
	if err != nil {
		fmt.Print(err)
	}
	user := &model.User{
		UserName:     data.UserName,
		HashPassword: data.HashPassword,
	}
	_, err = u.Service.AddUser(user)
	if err != nil {
		fmt.Print(err)
	}
	_, _ = Ctx.JSON(Success{
		"请求成功",
		"ok",
	})
}

func (u *UserController) PostLogin(Ctx iris.Context)  {
	var data model.User
	err := Ctx.ReadJSON(&data)
	if err != nil {
		fmt.Print(err)
	}
 	_, ok := u.Service.IsSuccess(data.UserName, data.HashPassword)
 	fmt.Printf("this is ok%v",ok)
 	if !ok {
 		_, _ = Ctx.JSON(Success{
 			"登陆失败",
 			"no",
		})
	}else {
		Ctx.SetCookieKV("test","1111")
		_, _ = Ctx.JSON(Success{
			"登陆成功",
			"ok",
		})
	}
}
