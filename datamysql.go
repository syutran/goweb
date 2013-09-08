//datamysql.go
package main

import (
	"database/sql"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"time"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang?charset=utf8")
	if err != nil {
		print("ERROR?")
		return
	}
	_, e2 := db.Query("select 1")
	if e2 == nil {
		println("DB OK")
	}
}
