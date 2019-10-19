package handle

import (
	myredis "WorkSpace/DevelopTest/file_store_server/cache/redis"
	mydb "WorkSpace/DevelopTest/file_store_server/db"
	"WorkSpace/DevelopTest/file_store_server/util"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io"
	"math"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type MultipartUploadInfo struct {
	FileHash   string
	FileSize   int
	UploadID   string
	ChunkSize  int
	ChunkCount int
}

func InitMultipartUploadHandle(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	userName := r.Form.Get("username")
	fileHash := r.Form.Get("filehash")
	fileSize, err := strconv.Atoi(r.Form.Get("filesize"))
	if err != nil {
		fmt.Println("upload trans err")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	conn := myredis.NewRedisClient().Get()
	defer conn.Close()

	upInfo := MultipartUploadInfo{
		FileHash:   fileHash,
		FileSize:   fileSize,
		UploadID:   userName + fmt.Sprintf("%x", time.Now().UnixNano()),
		ChunkSize:  5 * 1024 * 1024, //5M
		ChunkCount: int(math.Ceil(float64(fileSize) / (5 * 1024 * 1024))),
	}

	_, _ = conn.Do("HSET", "MP_"+upInfo.UploadID, "chunkcount", upInfo.ChunkCount)
	_, _ = conn.Do("HSET", "MP_"+upInfo.UploadID, "filehash", upInfo.FileHash)
	_, _ = conn.Do("HSET", "MP_"+upInfo.UploadID, "filesize", upInfo.FileSize)
	//写入分块信息
	_, _ = w.Write(util.NewResp(0, "OK", upInfo).JSONBytes())
}

//上传分块文件
func UploadPortHandle(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	uploadID := r.Form.Get("uploadid")
	chunkIndex := r.Form.Get("index")

	//创建目录
	fpath := "./file_store_server/data/" + uploadID + "/" + chunkIndex
	if err := os.MkdirAll(path.Dir(fpath), 0744); err != nil {
		fmt.Println("upload part create dir err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//创建文件
	fd, err := os.Create(fpath)
	if err != nil {
		fmt.Println("upload part create file err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	buf := make([]byte, 1024*1024)
	for {
		n, err := r.Body.Read(buf)
		if err != nil || err == io.EOF {
			break
		}
		_, _ = fd.Write(buf[:n])
	}

	conn := myredis.NewRedisClient().Get()
	defer conn.Close()

	_, _ = conn.Do("HSET", "MP_"+uploadID, "chkidx_"+chunkIndex, 1)
	_, _ = w.Write(util.NewResp(0, "OK", nil).JSONBytes())
}

func CompleteUploadHandle(w http.ResponseWriter, r *http.Request) {

	_ = r.ParseForm()
	upid := r.Form.Get("uploadid")
	userName := r.Form.Get("username")
	fileHash := r.Form.Get("filehash")
	fileName := r.Form.Get("filename")
	fileSize := r.Form.Get("filesize")

	conn := myredis.NewRedisClient().Get()
	defer conn.Close()

	data, err := redis.StringMap(conn.Do("HGETALL", "MP_"+upid))
	if err != nil {
		_, _ = w.Write(util.NewResp(-1, "complete upload filed", nil).JSONBytes())
		return
	}
	var (
		chunkCount int
		totalCount int
	)
	if countStr, ok := data["chunkcount"]; ok {
		t, err := strconv.Atoi(countStr)
		if err != nil {
			fmt.Println("strconv err: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		chunkCount = t
	}

	for key, value := range data {
		if strings.HasPrefix(key, "chkidx_") && value == "1" {
			totalCount++
		}
	}

	if chunkCount != totalCount {
		_, _ = w.Write(util.NewResp(-2, "invalid request", nil).JSONBytes())
		return
	}

	//TODO:分块合并
	fsize, _ := strconv.Atoi(fileSize)
	mydb.OnFileUploadFinished(fileHash, fileName, int64(fsize), "")
	mydb.OnUserFileUpdateFinished(userName, fileHash, fileName, int64(fsize))

	_, _ = w.Write(util.NewResp(0, "OK", nil).JSONBytes())
}
