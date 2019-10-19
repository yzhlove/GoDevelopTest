package db

import (
	mydb "WorkSpace/DevelopTest/file_store_server/db/mysql"
	"database/sql"
	"fmt"
	"time"
)

type UserFile struct {
	UserName     string
	FileHash     string
	FileName     string
	FileSize     int64
	UploadAt     string
	LastUpdateAt string
}

func OnUserFileUpdateFinished(userName, fileHash, fileName string, fileSize int64) bool {
	stmt, err := mydb.DBConn().Prepare("insert ignore into tbl_user_file (`user_name`,`file_sha1`,`file_name`" +
		",`file_size`,`upload_at`) values (?,?,?,?,?)")
	if err != nil {
		fmt.Println("file update finished err: ", err)
		return false
	}
	defer stmt.Close()
	if _, err := stmt.Exec(userName, fileHash, fileName, fileSize, time.Now()); err != nil {
		fmt.Println("file update finished exec err: ", err)
		return false
	}
	return true
}

func QueryFileMetas(userName string, limit int) (userFiles []UserFile, err error) {
	var stmt *sql.Stmt
	if stmt, err = mydb.DBConn().Prepare("select file_sha1,file_name,file_size,upload_at,last_update " +
		"from tbl_user_file where user_name=? limit ? "); err != nil {
		return
	}
	defer stmt.Close()
	var rows *sql.Rows
	if rows, err = stmt.Query(userName, limit); err != nil {
		return
	}
	userFiles = make([]UserFile, 0, 16)
	for rows.Next() {
		file := UserFile{}
		if err = rows.Scan(&file.FileHash, &file.FileName, &file.FileSize, &file.UploadAt, &file.LastUpdateAt);
			err != nil {
			return
		}
		userFiles = append(userFiles, file)
	}
	return
}
