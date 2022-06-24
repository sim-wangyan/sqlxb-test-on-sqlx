package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB
func InitSqlxDB() *sqlx.DB {

	var err interface{}
	Db, err = sqlx.Connect("mysql",
		"root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True")
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		if Db != nil  {
			Db.Close()
		}
	}
	Db.SetMaxOpenConns(16)
	Db.SetMaxIdleConns(8)

	return Db
}

func main()  {

	testBulderX()
}
