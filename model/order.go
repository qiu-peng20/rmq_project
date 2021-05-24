package model

type Order struct {
	ID int64 `json:"ID" sql:"ID"`
	UserId int64 `json:"user_id" sql:"userId"`
	ProductId int64 `json:"product_id" sql:"productId"`
	OrderStatus int64 `json:"order_status" sql:"orderStatus"`
}

const (
	OrderWait = iota
	OrderSuccess
	OrderFailed
)
