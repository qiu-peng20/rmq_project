package service

import (
	"RMQ_Project/DB"
	"RMQ_Project/model"
)

type ProductServiceInterface interface {
	GetProductByID(int64)(*model.Product,error)
	GetAllProduct()([]*model.Product,error)
	DeleteProductByID(int64)bool
	InsertProduct(product *model.Product)(int64,error)
	UpdateProduct(product *model.Product)error
}

type ProductService struct {
	DB.MyProduct
}

func NewProductService(p DB.MyProduct ) ProductServiceInterface {
	return &ProductService{p}
}

func (p ProductService)GetProductByID(id int64) (*model.Product, error)  {
	return p.SelectByKey(id)
}

func (p ProductService)GetAllProduct() ([]*model.Product,error)  {
	return p.SelectAll()
}

func (p ProductService)DeleteProductByID(id int64)bool  {
	return p.Delete(id)
}

func (p ProductService)InsertProduct(product *model.Product)(int64, error)  {
	return p.Inset(product)
}

func (p ProductService)UpdateProduct(product *model.Product)error  {
	return p.Update(product)
}
