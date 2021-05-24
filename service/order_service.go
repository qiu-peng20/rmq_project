package service

import (
	"RMQ_Project/DB"
	"RMQ_Project/model"
)

type OrderServiceInterface interface {
	InsertOrder(order *model.Order) (int64,error)
}

type OrderService struct {
	DB.OrderInterface
}

func NewOrderService(orderInterface DB.OrderInterface) OrderServiceInterface {
	return &OrderService{orderInterface}
}

func (o OrderService)InsertOrder(order *model.Order) (int64, error) {
	return o.Insert(order)
}
