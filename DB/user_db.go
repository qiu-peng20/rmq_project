package DB

import (
	"RMQ_Project/common"
	"RMQ_Project/model"
	"database/sql"
	"errors"
	"fmt"
)

type UserInterface interface {
	Conn() error
	Insert(user *model.User) (id int64, err error)
	Select(userName string) (*model.User, error)
}

type UserManager struct {
	table string
	db    *sql.DB
}

func NewUserInterface(table string, db *sql.DB) UserInterface {
	return &UserManager{table: table, db: db,}
}

func (u *UserManager) Conn() error {
	if u.db == nil {
		db, err := common.NewMysqlConn()
		if err != nil {
			fmt.Print(err)
			return err
		}
		u.db = db
	}
	if u.table == "" {
		u.table = "user"
	}
	return nil
}

func (u *UserManager) Select(userName string) (user *model.User, err error) {
	if userName == "" {
		return &model.User{}, errors.New("搜索条件不能为空")
	}
	if err := u.Conn(); err != nil {
		return &model.User{}, err
	}
	s := fmt.Sprintf("SELECT * FROM %v WHERE userName=?", u.table)
	rows, err := u.db.Query(s, userName)
	if err != nil {
		return &model.User{}, err
	}
	result := common.GetResultRow(rows)
	if len(result) == 0 {
		return &model.User{}, errors.New("用户不存在")
	}
	user = &model.User{}
	common.DataToStructByTagSql(result, user)
	return
}

func (u *UserManager) Insert(user *model.User) (id int64, err error) {
	if err = u.Conn(); err != nil {
		return 0, err
	}
	s := fmt.Sprintf("INSERT %v SET userName=?, hashPassword=?", u.table)
	stmt, err := u.db.Prepare(s)
	if err != nil {
		return id, err
	}
	result, err := stmt.Exec(user.UserName, user.HashPassword)
	if err != nil {
		return id, err
	}
	return result.LastInsertId()
}
