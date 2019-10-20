package handle

import (
	"WorkSpace/GoDevelopTest/file_store_service/account/proto"
	"context"
	"fmt"
)

type User struct{}

func (user *User) SignUp(ctx context.Context, req *proto.SignUpReq, resp *proto.SignUpResp) error {

	fmt.Println("UserName = ", req.UserName)
	fmt.Println("PassWord = ", req.PassWord)

	if req.UserName == "yzh" && req.PassWord == "123456" {
		resp.Code = 0
		resp.Message = "OK"
	} else {
		resp.Code = -1
		resp.Message = "FAIL"
	}
	return nil
}
