package db

import (
	mydb "WorkSpace/DevelopTest/file_store_server/db/mysql"
	"database/sql"
	"fmt"
)

type User struct {
	UserName     string `json:"Username"`
	Email        string `json:"Email"`
	Phone        string `json:"Phone"`
	SignUpAt     string `json:"SignupAt"`
	LastActiveAt string `json:"LastActiveAt"`
	Status       int    `json:"Status"`
}

//用户注册
func UserSignUp(name, passwd string) bool {
	stmt, err := mydb.DBConn().Prepare("insert ignore into tbl_user(`user_name`,`user_pwd`) values(?,?) ")
	if err != nil {
		fmt.Println("insert user err: ", err)
		return false
	}
	defer stmt.Close()
	ret, err := stmt.Exec(name, passwd)
	if err != nil {
		fmt.Println("exec user err: ", err)
		return false
	}
	if rowsAffected, err := ret.RowsAffected(); nil == err && rowsAffected > 0 {
		return true
	}
	return false
}

func UserSignIn(name, encpasswd string) bool {
	stmt, err := mydb.DBConn().Prepare("select * from tbl_user where user_name=? limit 1")
	if err != nil {
		fmt.Println("query user err:", err)
		return false
	}
	defer stmt.Close()
	rows, err := stmt.Query(name)
	if err != nil {
		fmt.Println("stmt query user err: ", err)
		return false
	}
	if rows == nil {
		fmt.Println("query  not found name : ", name)
		return false
	}
	mapRows := mydb.ToMap(rows)
	if len(mapRows) > 0 && string(mapRows[0]["user_pwd"].([]byte)) == encpasswd {
		return true
	}
	return false
}

func UpdateToken(name, token string) bool {
	stmt, err := mydb.DBConn().Prepare("replace into tbl_user_token(`user_name`,`user_token`) values (?,?)")
	if err != nil {
		fmt.Println("update token prepar err: ", err)
		return false
	}
	defer stmt.Close()
	if _, err = stmt.Exec(name, token); err != nil {
		fmt.Println("update token exec err: ", err)
		return false
	}
	return true
}

func GetUserInfo(name string) (user User, err error) {
	var stmt *sql.Stmt
	if stmt, err = mydb.DBConn().Prepare("select user_name,signup_at from tbl_user where user_name=? limit 1"); err != nil {
		fmt.Println("get user info err: ", err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(name).Scan(&user.UserName, &user.SignUpAt)
	return
}

func GetToken(name string) (token string) {
	stmt, err := mydb.DBConn().Prepare("select user_token from tbl_user_token where user_name=? limit 1")
	if err != nil {
		fmt.Println("get token err: ", err)
		return
	}
	defer stmt.Close()
	if err = stmt.QueryRow(name).Scan(&token); err != nil {
		return ""
	}
	return
}
