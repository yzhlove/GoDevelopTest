package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db *sql.DB

func init() {
	var err error
	if db, err = sql.Open("mysql",
		"root:123456@tcp(localhost:3306)/fileserver?charset=utf8"); err != nil {
		fmt.Println("connect mysql err:", err)
		return
	}
	db.SetMaxOpenConns(1000)
	if err = db.Ping(); err != nil {
		fmt.Println("ping mysql err:", err)
		os.Exit(1)
	}
}

//返回数据库链接对象
func DBConn() *sql.DB {
	return db
}

func ToMap(rows *sql.Rows) (data []map[string]interface{}) {
	columns, err := rows.Columns()
	if err != nil {
		return
	}
	args := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		args[i] = &values[i]
	}
	table := make(map[string]interface{})
	data = make([]map[string]interface{}, 0)
	for rows.Next() {
		isErr(rows.Scan(args...))
		for i, col := range values {
			if col != nil {
				table[columns[i]] = col
			}
		}
		data = append(data, table)
	}
	return
}

func isErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
