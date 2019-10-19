package db

import (
	mydb "WorkSpace/DevelopTest/file_store_server/db/mysql"
	"database/sql"
	"fmt"
)

type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

func OnFileUploadFinished(hash, name string, size int64, addr string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_file (`file_sha1`,`file_name`,`file_size`,`file_addr`,`status`) " +
			"values(?,?,?,?,1)")
	if err != nil {
		fmt.Println("paper statement err : ", err)
		return false
	}
	ret, err := stmt.Exec(hash, name, size, addr)
	if err != nil {
		fmt.Println("exec err :", err)
		return false
	}
	if rf, err := ret.RowsAffected(); nil == err {
		if rf <= 0 {
			fmt.Printf("%s has been upload before", hash)
		}
		return true
	}
	return false
}

func GetFileMeta(hash string) (status *TableFile, err error) {
	var stmt *sql.Stmt
	if stmt, err = mydb.DBConn().Prepare(
		"select file_sha1,file_name,file_size,file_addr from tbl_file " +
			"where file_sha1=? and status=1 limit 1"); err != nil {
		return
	}
	defer stmt.Close()
	status = new(TableFile)
	if err = stmt.QueryRow(hash).
		Scan(&status.FileHash, &status.FileName, &status.FileSize, &status.FileAddr); err != nil {
		return
	}
	return
}
