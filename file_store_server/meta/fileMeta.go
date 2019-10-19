package meta

import (
	mydb "WorkSpace/GoDevelopTest/file_store_server/db"
	"fmt"
)

type SHA1 string

//文件元信息结构
type FileMeta struct {
	FileSha1   SHA1
	FileName   string
	FileSize   int64
	Local      string
	UploadText string
}

var (
	fileMetas map[SHA1]FileMeta
	initMax   = 16
)

func init() {
	fileMetas = make(map[SHA1]FileMeta, initMax)
}

//更新
func UpdateFileMetas(fm FileMeta) {
	fileMetas[fm.FileSha1] = fm
}

func UpdateFileMetaDB(fm FileMeta) bool {
	return mydb.OnFileUploadFinished(
		string(fm.FileSha1), fm.FileName, fm.FileSize, fm.Local)
}

//获取
func GetFileMeta(sha1 SHA1) FileMeta {
	return fileMetas[sha1]
}

func GetFilemetaDB(sha1 SHA1) (FileMeta, error) {
	status, err := mydb.GetFileMeta(string(sha1))
	if err != nil {
		fmt.Println("GetFileMeta Err: ", err)
		return FileMeta{}, err
	}
	fm := FileMeta{
		FileSha1: SHA1(status.FileHash),
		FileName: status.FileName.String,
		FileSize: status.FileSize.Int64,
		Local:    status.FileAddr.String,
	}
	return fm, nil
}

//删除
func RemoveFileMeta(sha1 SHA1) {
	delete(fileMetas, sha1)
}
