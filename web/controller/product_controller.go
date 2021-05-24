package controller

import (
	"RMQ_Project/common"
	"RMQ_Project/model"
	"RMQ_Project/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ProductController struct {
	service.ProductServiceInterface
}

func (p *ProductController) GetAll(Ctx iris.Context) ([]*model.Product,error)  {
	return  p.GetAllProduct()
}

func (p *ProductController) GetAdd(Ctx iris.Context) mvc.View {
	return mvc.View{
		Name: "product/add.html",
	}
}

func (p *ProductController) PostAdd(Ctx iris.Context) {
	product := &model.Product{}
	_ = Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{"imooc", true, true})
	if err := dec.Decode(Ctx.Request().Form, product); err != nil {
		Ctx.Application().Logger().Debug(err)
	}
	_, err := p.ProductServiceInterface.InsertProduct(product)
	if err != nil {
		Ctx.Application().Logger().Debug(err)
	}
	Ctx.Redirect("/product/all")
}
