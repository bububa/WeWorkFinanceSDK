package wxworkfinancesdk

import (
	"fmt"
)

const (
	// SDKErrMsg 错误消息
	SDKErrMsg = "sdk failed"
)

// Error SDK Error结构体
type Error struct {
	ErrCode int    `json:"errcode,omitempty"` // 错误代码
	ErrMsg  string `json:"errmsg,omitempty"`  // 错误说明
}

// Error implement error interface
func (e Error) Error() string {
	return fmt.Sprintf("%d:%s", e.ErrCode, e.ErrMsg)
}

// NewSDKErr 新建错误
func NewSDKErr(code int) Error {
	return Error{ErrCode: code, ErrMsg: SDKErrMsg}
}
