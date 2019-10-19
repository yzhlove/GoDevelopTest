package handle

import (
	mydb "WorkSpace/DevelopTest/file_store_server/db"
	"WorkSpace/DevelopTest/file_store_server/meta"
	"WorkSpace/DevelopTest/file_store_server/util"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func UploadHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//返回上传的Html页面
		data, err := ioutil.ReadFile("./file_store_server/static/view/index.html")
		if err != nil {
			_, _ = io.WriteString(w, "internel server error")
			return
		}
		_, _ = io.WriteString(w, string(data))
		return
	} else if r.Method == http.MethodPost {
		userName := "admin"
		//接受文件流存储在本地
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data , err : %v \n", err)
			return
		}
		defer file.Close()

		//创建文件元信息
		fm := meta.FileMeta{
			FileName:   head.Filename,
			Local:      "./file_store_server/tmp/" + head.Filename,
			UploadText: time.Now().Format("2006-01-02 15:04:04"),
		}

		newFile, err := os.Create(fm.Local)
		if err != nil {
			fmt.Printf("create file err : %v \n", err)
			return
		}
		defer newFile.Close()

		fm.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("copy file err : %v \n", err)
			return
		}
		_, _ = newFile.Seek(0, 0)
		fm.FileSha1 = meta.SHA1(util.FileSha1(newFile))

		_ = meta.UpdateFileMetaDB(fm)

		if !mydb.OnUserFileUpdateFinished(userName, string(fm.FileSha1), fm.FileName, fm.FileSize) {
			fmt.Println("upload user file err")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		//302 临时重定向
		http.Redirect(w, r, "http://"+r.Host+"/static/view/home.html", http.StatusFound)
	}
}

func GetFileMetaHandle(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fileHash := r.Form["file_hash"][0]
	fm, err := meta.GetFilemetaDB(meta.SHA1(fileHash))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(fm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(data)
}

func FileQueryHandle(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	userName := r.Form.Get("username")
	limit := r.Form.Get("limit")
	cnt, err := strconv.Atoi(limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userFiles, err := mydb.QueryFileMetas(userName, cnt)
	if err != nil {
		fmt.Println("file query err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(userFiles)
	if err != nil {
		fmt.Println("query file json err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(data)
}

func DownloadHandle(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fileHash := r.Form["file_hash"][0]

	fm := meta.GetFileMeta(meta.SHA1(fileHash))
	f, err := os.Open(fm.Local)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("Content-Disposition", "attachment;filename=\"love"+fm.FileName+"\"")
	_, _ = w.Write(data)

}

func FileMetaUpdateHandle(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_ = r.ParseForm()
	fileHash := r.Form.Get("file_hash")
	newFileName := r.Form.Get("file_name")

	fm := meta.GetFileMeta(meta.SHA1(fileHash))
	fm.FileName = newFileName
	meta.UpdateFileMetas(fm)

	data, err := json.Marshal(fm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func FileDeleteHandle(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fileHash := r.Form.Get("file_hash")
	fm := meta.GetFileMeta(meta.SHA1(fileHash))
	_ = os.Remove(fm.Local)
	meta.RemoveFileMeta(meta.SHA1(fileHash))
	w.WriteHeader(http.StatusOK)
}

func TryFastUploadHandle(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	userName := r.Form.Get("username")
	fileHash := r.Form.Get("filehash")
	fileName := r.Form.Get("filename")
	fileSizeStr := r.Form.Get("file")
	fileSize, err := strconv.Atoi(fileSizeStr)
	if err != nil {
		fmt.Println("trans int err: ", err)
		return
	}
	fm, err := meta.GetFilemetaDB(meta.SHA1(fileHash))
	if err != nil {
		fmt.Println("not found file: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if "" == fm.FileSha1 {
		resp := util.NewResp(-1, "try fast err", nil)
		_, _ = w.Write(resp.JSONBytes())
		return
	}
	var respMsg util.RespMsg
	if !mydb.OnUserFileUpdateFinished(userName, fileHash, fileName, int64(fileSize)) {
		respMsg = util.RespMsg{
			Code: 0,
			Msg:  "try fast OK",
			Data: nil,
		}
	} else {
		respMsg = util.RespMsg{
			Code: -2,
			Msg:  "try fast again",
			Data: nil,
		}
	}
	_, _ = w.Write(respMsg.JSONBytes())
}
