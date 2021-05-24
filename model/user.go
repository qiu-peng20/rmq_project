package model

type User struct {
	ID int64 `json:"ID" sql:"ID" form:"ID"`
	UserName string `json:"userName" sql:"userName" form:"userName"`
	HashPassword string `json:"hash_password" sql:"hashPassword" form:"hashPassword" `
}
