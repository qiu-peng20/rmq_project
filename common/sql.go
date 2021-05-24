package common

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql")

//创建数据库链接
func NewMysqlConn() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/rabbitmq?charset=utf8")
	return
}

// 获取一条返回值
func GetResult()  {
	
}