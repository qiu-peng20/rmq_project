package DB

import (
	"RMQ_Project/common"
	"RMQ_Project/model"
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

type MyProduct interface {
	Conn() error
	Inset(*model.Product) (int64, error)
	Delete(int64) bool
	Update(*model.Product) error
	SelectByKey(int64) (*model.Product, error)
	SelectAll() ([]*model.Product, error)
}

type ProductManager struct {
	table string
	conn  *sql.DB
}

func NewProductManager(table string, conn *sql.DB) MyProduct {
	return &ProductManager{table: table, conn: conn}
}

// 链接数据库
func (pm ProductManager) Conn() (err error) {
	if pm.conn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		pm.conn = mysql
	}
	if pm.table == "" {
		pm.table = "product"
	}
	return
}

//插入数据
func (pm ProductManager) Inset(product *model.Product) (productID int64, err error) {
	if err := pm.Conn(); err != nil {
		fmt.Print(err)
		return 0, err
	}
	// 准备sql
	const SQL = `INSERT product SET productName=?, productNum=?, productImage=?, productUrl=?`
	stmt, err := pm.conn.Prepare(SQL)
	if err != nil {
		fmt.Print(err)
		return 0, err
	}
	defer stmt.Close()
	//插入数据
	result, err := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		fmt.Print(err)
		return 0, err
	}
	//返回数据
	fmt.Print(result.LastInsertId())
	return result.LastInsertId()
}

func (pm ProductManager) Delete(id int64) bool {
	if err := pm.Conn(); err == nil {
		return false
	}
	const SQL = `DELETE FROM product WHERE id=?`
	stmt, err := pm.conn.Prepare(SQL)
	if err != nil {
		log.Print(err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(strconv.FormatInt(id, 10))
	if err != nil {
		log.Print(err)
		return false
	}
	return true
}

func (pm ProductManager) Update(product *model.Product) error {
	if err := pm.Conn(); err == nil {
		return err
	}
	const SQL = `UPDATE product SET productName=?, productNum=?, productImage=?, productUrl=? WHERE id =?`
	smtp, err := pm.conn.Prepare(SQL)
	if err != nil {
		return err
	}
	defer smtp.Close()
	_, err = smtp.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func (pm ProductManager) SelectByKey(id int64) (productResult *model.Product, err error) {
	if err := pm.Conn(); err == nil {
		return &model.Product{}, err
	}
	SQL := "SELECT * FROM " + pm.table + "WHERE id=?"
	row, err := pm.conn.Query(SQL)
	if err != nil {
		return &model.Product{}, err
	}
	defer row.Close()
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &model.Product{}, nil
	}
	common.DataToStructByTagSql(result, &model.Product{})
	return
}

func (pm ProductManager) SelectAll() (productArray []*model.Product, err error) {
	if err := pm.Conn(); err != nil {
		return nil, err
	}
	SQL := fmt.Sprintf("SELECT * FROM %s",pm.table)
	rows, err := pm.conn.Query(SQL)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	defer rows.Close()
	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return nil, nil
	}
	for _, v := range result {
		product := &model.Product{}
		common.DataToStructByTagSql(v, product)
		productArray = append(productArray, product)
	}
	return productArray,err
}
