package util

import (
	"encoding/json"
	"fmt"
)

type RespMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResp(code int, msg string, data interface{}) *RespMsg {
	return &RespMsg{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func (resp *RespMsg) JSONBytes() []byte {
	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("JsonMarshal Err: %v \n", err)
		return nil
	}
	return data
}

func (resp *RespMsg) ToSimpleJSONBytes(code int, msg string) []byte {
	return []byte(fmt.Sprintf(`{"code":%d,"msg":%s}`, code, msg))
}

func (resp *RespMsg) ToSimpleJSON(code int, msg string) string {
	return fmt.Sprintf(`{"code":%d,"msg":%s}`, code, msg)
}
