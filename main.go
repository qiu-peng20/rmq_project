package main

import (
	"RMQ_Project/DB"
	"RMQ_Project/common"
	"RMQ_Project/service"
	"RMQ_Project/web/controller"
	"fmt"
	"github.com/kataras/iris/v12"
)

func main() {
	// 1 创建iris实例
	app := iris.New()
	// 2 设置错误模式
	app.Logger().SetLevel("debug")
	// 3 注册模板
	template := iris.HTML("./web/view", ".html").Layout(
		"shared/layout.html").Reload(true)
	app.RegisterView(template)
	// 4 设置模板目标
	app.HandleDir("/assets","./web/assets")
	// 注册数据库
	db, err := common.NewMysqlConn()
	if err != nil {
		fmt.Print(err)
	}
	// 注册上下文
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	//5 注册控制器
	productRepository := DB.NewProductManager("product", db)
	productService := service.NewProductService(productRepository)
	p  := new(controller.ProductController)
	p.ProductServiceInterface = productService
	app.Post("/product", p.PostAdd)
	//productParty := app.Party("/product")
	//product := mvc.New(productParty)
	//product.Register(ctx, productService)
	//product.Handle(new(controller.ProductController))

	// 注册用户控制器
	userRepository := DB.NewUserInterface("user", db)
	userInterface := service.NewUserService(userRepository)
	u := new(controller.UserController)
	u.Service = userInterface
	app.Post("/user/sign", u.PostRegister)
	app.Post("/user/login", u.PostLogin)
	
	//6 启动服务
	_ = app.Run(iris.Addr("localhost:8080"))
}
