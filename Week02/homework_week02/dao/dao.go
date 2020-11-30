package dao

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	myerrors "github.com/pkg/errors"
	"homework_week02/model"
)

const dataSource = "root:root@tcp(127.0.0.1:3306)/test?charset=utf8"

var (
	UserNotFoundErr = errors.New("user not found")
)

type Dao struct {
}

var DB *sql.DB

func init() {
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	DB = db
}

func (d *Dao) GetUserInfoById(id int64) (*model.UserInfo, error) {
	var user model.UserInfo
	err := DB.QueryRow("select `id`, `name`, `age` from student where id=?", id).Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, UserNotFoundErr
		}
		return &user, myerrors.Wrap(err, "query user failed")
	}
	return &user, nil
}
