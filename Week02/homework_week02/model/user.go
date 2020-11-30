package model

type UserInfo struct {
	Id int64 `db:"uid"`
	Name string `db:"userName"`
	Age uint `db:"userAge"`
}
