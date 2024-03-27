package generator

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlConnector struct {
	DatabaseName string
	Username     string
	Password     string
	Host         string
	Port         string

	db *sql.DB
}

func (mc *MysqlConnector) connect() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mc.Username, mc.Password, mc.Host, mc.Port, mc.DatabaseName))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Success!")

	mc.db = db
}

func (mc *MysqlConnector) close() {
	if mc.db != nil {
		mc.db.Close()
	}
}
