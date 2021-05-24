package model

type User struct {
	ID int64 `json:"ID" sql:"ID" form:"ID"`
	UserName string `json:"userName" sql:"userName" form:"userName"`
	HashPassword string `json:"hashPassword" sql:"hashPassword" form:"hashPassword" `
}
