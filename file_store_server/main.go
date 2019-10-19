package main

import (
	"WorkSpace/GoDevelopTest/file_store_server/handle"
	"fmt"
	"net/http"
)

const (
	static = "/Users/yurisa/Develop/GoWork/src/WorkSpace/DevelopTest/file_store_server/static"
)

func main() {
	fmt.Println("start file-store-server ....")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(static))))
	http.HandleFunc("/file/upload", handle.UploadHandle)
	http.HandleFunc("/file/meta", handle.GetFileMetaHandle)
	http.HandleFunc("/file/download", handle.DownloadHandle)
	http.HandleFunc("/file/update", handle.FileMetaUpdateHandle)
	http.HandleFunc("/file/delete", handle.FileDeleteHandle)
	http.HandleFunc("/file/query", handle.FileQueryHandle)
	http.HandleFunc("/file/fast_upload", handle.TryFastUploadHandle)

	http.HandleFunc("/user/signup", handle.SignUpHandle)
	http.HandleFunc("/user/signin", handle.SignInHandle)
	http.HandleFunc("/user/info", handle.HttpInterceptor(handle.UserInfoHandle))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("start err:%v\n", err)
	}
}
