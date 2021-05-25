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
	if err = u.Conn(); err != nil {
		return &model.User{}, err
	}
	s := fmt.Sprintf("SELECT * FROM %v WHERE userName=? LIMIT 1", u.table)
	stmt, err := u.db.Prepare(s)
	if err != nil {
		return &model.User{}, err
	}
	defer stmt.Close()
	user = &model.User{}
	err = stmt.QueryRow(userName).Scan(&user.ID, &user.UserName, &user.HashPassword)
	if err != nil {
		fmt.Printf("this is scan error %v",err)
		return user, err
	}
	if user.HashPassword == "" {
		fmt.Print("没有这个用户")
		return user, err
	}
	return user, nil
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
