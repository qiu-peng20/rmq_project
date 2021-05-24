package DB

import (
	"RMQ_Project/common"
	"RMQ_Project/model"
	"database/sql"
	"fmt"
)

type OrderInterface interface {
	Conn() error
	Insert(order *model.Order) (int64, error)
	SelectByAll() ([]*model.Order, error)
}

type OrderStruct struct {
	table string
	db    *sql.DB
}

func NewOrderManger(table string, sql *sql.DB) OrderInterface {
	return &OrderStruct{
		table: table,
		db:    sql,
	}
}

func (o *OrderStruct) Conn() (err error) {
	if o.db == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			fmt.Print(err)
			return err
		}
		o.db = mysql
	}
	if o.table == "" {
		o.table = "order"
	}
	return nil
}

func (o OrderStruct) Insert(order *model.Order) (id int64, err error) {
	if err := o.Conn(); err != nil {
		fmt.Print(err)
		return
	}
	s := fmt.Sprintf("INSERT %v SET userId=?, productId=?, orderStatus=?",o.table)
	stmt, err := o.db.Prepare(s)
	if err != nil {
		return id,err
	}
	result, err := stmt.Exec(order.UserId, order.ProductId, order.OrderStatus)
	if err != nil {
		fmt.Print(err)
		return id,err
	}
	return result.LastInsertId()
}

func (o OrderStruct) SelectByAll() ([]*model.Order, error) {

}
